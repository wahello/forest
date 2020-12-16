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

package helper

import (
	"errors"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/tidwall/gjson"
	"github.com/vicanso/elton"
	"github.com/vicanso/forest/cs"
	"github.com/vicanso/forest/util"
	"github.com/vicanso/go-axios"
	"github.com/vicanso/hes"
	"go.uber.org/zap"
)

const (
	httpErrCategoryDNS     = "dns"
	httpErrCategoryTimeout = "timeout"
	httpErrCategoryAddr    = "addr"
	httpErrCategoryAborted = "aborted"
	httpErrCategoryRefused = "refused"
	httpErrCategoryReset   = "reset"
)

func newHTTPOnDone(serviceName string) axios.OnDone {
	return func(conf *axios.Config, resp *axios.Response, err error) {
		ht := conf.HTTPTrace

		reused := false
		addr := ""
		use := ""
		ms := 0
		id := conf.GetString(cs.CID)
		status := -1
		if ht != nil {
			reused = ht.Reused
			addr = ht.Addr
			timelineStats := ht.Stats()
			use = timelineStats.String()
			ms = int(timelineStats.Total.Milliseconds())
		}
		if resp != nil {
			status = conf.Response.Status
		}

		tags := map[string]string{
			"service": serviceName,
			"route":   conf.Route,
			"method":  conf.Method,
		}
		fields := map[string]interface{}{
			"cid":    id,
			"url":    conf.URL,
			"status": status,
			"addr":   addr,
			"reused": reused,
			"use":    ms,
		}
		message := ""
		if err != nil {
			message = err.Error()
			fields["error"] = message
			errCategory := getHTTPErrorCategory(err)
			if errCategory != "" {
				fields["errCategory"] = errCategory
			}
		}
		logger.Info("http request stats",
			zap.String("service", serviceName),
			zap.String("cid", id),
			zap.String("method", conf.Method),
			zap.String("route", conf.Route),
			zap.String("url", conf.URL),
			zap.Any("params", conf.Params),
			zap.Any("query", conf.Query),
			zap.Int("status", status),
			zap.String("addr", addr),
			zap.Bool("reused", reused),
			zap.String("use", use),
			zap.String("error", message),
		)
		GetInfluxSrv().Write(cs.MeasurementHTTPRequest, tags, fields)

	}
}

// newHTTPConvertResponseToError 将http响应码为>=400的转换为出错
func newHTTPConvertResponseToError(serviceName string) axios.ResponseInterceptor {
	return func(resp *axios.Response) (err error) {
		if resp.Status >= 400 {
			message := gjson.GetBytes(resp.Data, "message").String()
			if message == "" {
				message = string(resp.Data)
			}
			// 只返回普通的error对象，由onError来转换为http error
			err = errors.New(message)
		}
		return
	}
}

// getHTTPErrorCategory 获取出错的类型，主要分类DNS错误，addr错误以及一些系统调用的异常
func getHTTPErrorCategory(err error) string {

	netErr, ok := err.(net.Error)
	if ok && netErr.Timeout() {
		return httpErrCategoryTimeout
	}

	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return httpErrCategoryDNS
	}
	var addrErr *net.AddrError
	if errors.As(err, &addrErr) {
		return httpErrCategoryAddr
	}

	opErr, ok := netErr.(*net.OpError)
	if !ok {
		return ""
	}
	switch e := opErr.Err.(type) {
	// 针对以下几种系统调用返回对应类型
	case *os.SyscallError:
		if no, ok := e.Err.(syscall.Errno); ok {
			switch no {
			case syscall.ECONNREFUSED:
				return httpErrCategoryRefused
			case syscall.ECONNABORTED:
				return httpErrCategoryAborted
			case syscall.ECONNRESET:
				return httpErrCategoryReset
			case syscall.ETIMEDOUT:
				return httpErrCategoryTimeout
			}
		}
	}

	return ""
}

// newHTTPOnError 新建error的处理函数
func newHTTPOnError(serviceName string) axios.OnError {
	return func(err error, conf *axios.Config) (newErr error) {
		code := -1
		if conf.Response != nil {
			code = conf.Response.Status
		}
		he := hes.Wrap(err)
		if code >= http.StatusBadRequest {
			he.StatusCode = code
		}
		// 如果未设置http响应码，则设置为500
		if he.StatusCode < http.StatusBadRequest {
			he.StatusCode = http.StatusInternalServerError
		}

		if he.Extra == nil {
			he.Extra = make(map[string]interface{})
		}

		he.Category = getHTTPErrorCategory(err)

		if !util.IsProduction() {
			he.Extra["requestRoute"] = conf.Route
			he.Extra["requestService"] = serviceName
			he.Extra["requestCURL"] = conf.CURL()
			// TODO 是否非生产环境增加更多的信息，方便测试时确认问题
		}
		return he
	}
}

// NewHTTPInstance 新建实例
func NewHTTPInstance(serviceName, baseURL string, timeout time.Duration) *axios.Instance {
	return axios.NewInstance(&axios.InstanceConfig{
		EnableTrace: true,
		Timeout:     timeout,
		OnError:     newHTTPOnError(serviceName),
		OnDone:      newHTTPOnDone(serviceName),
		BaseURL:     baseURL,
		ResponseInterceptors: []axios.ResponseInterceptor{
			newHTTPConvertResponseToError(serviceName),
		},
	})
}

// AttachWithContext 添加context中的cid至请求的config中
func AttachWithContext(conf *axios.Config, c *elton.Context) {
	if c == nil || conf == nil {
		return
	}
	conf.Set(cs.CID, c.ID)
	if conf.Context == nil {
		conf.Context = c.Context()
	}
}
