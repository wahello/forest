// Copyright 2020 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/vicanso/elton"
	compress "github.com/vicanso/elton-compress"
	M "github.com/vicanso/elton/middleware"
	"github.com/vicanso/forest/config"
	_ "github.com/vicanso/forest/controller"
	"github.com/vicanso/forest/helper"
	"github.com/vicanso/forest/log"
	"github.com/vicanso/forest/middleware"
	"github.com/vicanso/forest/router"
	"github.com/vicanso/forest/service"
	"github.com/vicanso/forest/util"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// 是否用户主动关闭
var closedByUser = false

// watchForClose 监听信号关闭程序
func watchForClose(e *elton.Elton, fn func()) {
	logger := log.Default()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			logger.Info("server will be closed",
				zap.String("signal", s.String()),
			)
			closedByUser = true
			// 设置状态为退出中，/ping请求返回出错，反向代理不再转发流量
			config.SetApplicationStatus(config.ApplicationStatusStopping)
			// docker 在10秒内退出，因此设置8秒后退出
			time.Sleep(5 * time.Second)
			// 所有新的请求均返回出错
			e.GracefulClose(3 * time.Second)
			fn()
			os.Exit(0)
		}
	}()
}

// 相关依赖服务的校验，主要是数据库等
func dependServiceCheck() (err error) {
	err = helper.RedisPing()
	if err != nil {
		return
	}
	// configSrv := new(service.ConfigurationSrv)
	// err = configSrv.Refresh()
	// if err != nil {
	// 	return
	// }
	return
}

func main() {
	// TODO defer panic
	logger := log.Default()
	closeOnce := sync.Once{}
	closeDeps := func() {
		closeOnce.Do(func() {
			// 关闭influxdb，flush统计数据
			helper.GetInfluxSrv().Close()
		})
	}
	defer closeDeps()
	basicConfig := config.GetBasicConfig()

	e := elton.New()
	// 启用耗时跟踪
	if util.IsDevelopment() {
		e.EnableTrace = true
	}
	e.SignedKeys = service.GetSignedKeys()
	e.OnTrace(func(c *elton.Context, infos elton.TraceInfos) {
		// 设置server timing
		c.ServerTiming(infos, "forest-")
	})

	// 非开发环境，监听信号退出
	if !util.IsDevelopment() {
		watchForClose(e, closeDeps)
	}

	// 捕捉panic异常，避免程序崩溃
	e.UseWithName(M.NewRecover(), "recover")

	// 入口设置
	e.UseWithName(middleware.NewEntry(), "entry")

	// 接口相关统计信息
	e.UseWithName(middleware.NewStats(), "stats")

	// 出错转换为json（出错处理应该在stats之后，这样stats中才可获取到正确的http status code)
	e.UseWithName(middleware.NewError(), "error")

	// 限制最大请求量
	if basicConfig.RequestLimit != 0 {
		e.UseWithName(M.NewGlobalConcurrentLimiter(M.GlobalConcurrentLimiterConfig{
			Max: uint32(basicConfig.RequestLimit),
		}), "requestLimit")
	}

	// 配置只针对snappy与lz4压缩（主要用于减少内网线路带宽，对外的压缩由前置反向代理 完成）
	compressMinLength := 2 * 1024
	compressConfig := M.NewCompressConfig(
		&compress.SnappyCompressor{
			MinLength: compressMinLength,
		},
		&compress.Lz4Compressor{
			MinLength: compressMinLength,
		},
	)
	e.UseWithName(M.NewCompress(compressConfig), "compress")

	// IP限制
	e.UseWithName(middleware.NewIPBlocker(), "ipBlocker")

	// 根据配置对路由mock返回
	e.UseWithName(middleware.NewRouterMocker(), "routerMocker")

	// 路由并发限制
	e.UseWithName(M.NewRCL(M.RCLConfig{
		Limiter: service.GetRouterConcurrencyLimiter(),
	}), "rcl")

	// etag与fresh的处理
	e.UseWithName(M.NewDefaultFresh(), "fresh").
		UseWithName(M.NewDefaultETag(), "etag")

	// 对响应数据 c.Body 转换为相应的json响应
	e.UseWithName(M.NewDefaultResponder(), "responder")

	// 读取读取body的数的，转换为json bytes
	e.UseWithName(M.NewDefaultBodyParser(), "body-parser")

	// 初始化路由
	for _, g := range router.GetGroups() {
		e.AddGroup(g)
	}

	// 初始化路由并发限制配置
	service.InitRouterConcurrencyLimiter(e.Routers)

	err := dependServiceCheck()
	if err != nil {
		service.AlarmError("check depend service fail, " + err.Error())
		logger.DPanic("exception",
			zap.Error(err),
		)
		return
	}

	config.SetApplicationStatus(config.ApplicationStatusRunning)

	// http1与http2均支持
	e.Server = &http.Server{
		Handler: h2c.NewHandler(e, &http2.Server{}),
	}
	logger.Info("server will listen on " + basicConfig.Listen)
	err = e.ListenAndServe(basicConfig.Listen)
	// 如果出错而且非用户关闭，则发送告警
	if err != nil && !closedByUser {
		service.AlarmError("listen and serve fail, " + err.Error())
		logger.DPanic("exception",
			zap.Error(err),
		)
	}
}
