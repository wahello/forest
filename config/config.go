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

// 应用中的所有配置获取，拉取配置信息使用default配置中的值为默认值，再根据GO_ENV配置的环境变量获取对应的环境配置

package config

import (
	"bytes"
	"io"
	"net/url"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/vicanso/forest/validate"
	"github.com/vicanso/viperx"
)

var (
	box = packr.New("config", "../configs")
	env = os.Getenv("GO_ENV")

	defaultViperX *viperx.ViperX

	// 应用状态
	applicationStatus = ApplicationStatusStopped

	// 应用名称
	appName string
	// 应用版本
	version string
	// 应用构建时间
	buildedAt string
)

const (
	// Dev 开发模式下的环境变量
	Dev = "dev"
	// Test 测试环境下的环境变量
	Test = "test"
	// Production 生产环境下的环境变量
	Production = "production"
)

const (
	// ApplicationStatusStopped 应用停止
	ApplicationStatusStopped int32 = iota
	// ApplicationStatusRunning 应用运行中
	ApplicationStatusRunning
	// ApplicationStatusStopping 应用正在停止
	ApplicationStatusStopping
)

type (
	// BasicConfig 应用基本配置信息
	BasicConfig struct {
		// 监听地址
		Listen string `validate:"ascii,required"`
		// 最大处理请求数
		RequestLimit uint `validate:"required"`
		// 应用名称
		Name string `validate:"ascii"`
	}
	// SessionConfig session相关配置信息
	SessionConfig struct {
		// cookie的保存路径
		CookiePath string `validate:"ascii,required"`
		// cookie的key
		Key string `validate:"ascii,required"`
		// cookie的有效期
		TTL time.Duration `validate:"required"`
		// 用于加密cookie的key
		Keys []string `validate:"required"`
		// 用于跟踪用户的cookie
		TrackKey string `validate:"ascii,required"`
	}
	// RedisConfig redis配置
	RedisConfig struct {
		// 连接地址
		Addr string `validate:"hostname_port,required"`
		// 密码
		Password string
		// db(0,1,2等)
		DB int
		// 慢请求时长
		Slow time.Duration `validate:"required"`
		// 最大的正在处理请求量
		MaxProcessing uint32 `validate:"required"`
	}
	// MailConfig email的配置
	MailConfig struct {
		Host     string `validate:"hostname,required"`
		Port     int    `validate:"number,required"`
		User     string `validate:"email,required"`
		Password string `validate:"min=1,max=100"`
	}
	// Influxdb influxdb配置
	InfluxdbConfig struct {
		// 存储的bucket
		Bucket string `validate:"min=1,max=50"`
		// 配置的组织名称
		Org string `validate:"min=1,max=100"`
		// 连接地址
		URI string `validate:"url,required"`
		// 认证的token
		Token string `validate:"ascii,required"`
		// 批量提交大小
		BatchSize uint `validate:"min=1,max=5000"`
		// 间隔提交时长
		FlushInterval time.Duration `validate:"required"`
		// 是否禁用
		Disabled bool
	}
	// AlarmConfig alarm配置
	AlarmConfig struct {
		// 接收人列表
		Receivers []string `validate:"required"`
	}
)

func init() {

	configType := "yml"
	defaultViperX = viperx.New(configType)

	configExt := "." + configType
	readers := make([]io.Reader, 0)
	for _, name := range []string{
		"default",
		GetENV(),
	} {
		data, err := box.Find(name + configExt)
		if err != nil {
			panic(err)
		}
		readers = append(readers, bytes.NewReader(data))
	}

	err := defaultViperX.ReadConfig(readers...)
	if err != nil {
		panic(err)
	}

	// appName = GetString("app")
}

func validatePanic(v interface{}) {
	err := validate.Do(v, nil)
	if err != nil {
		panic(err)
	}
}

// GetENV 获取当前运行环境
func GetENV() string {
	if env == "" {
		return Dev
	}
	return env
}

// SetApplicationStatus 设置应用运行状态
func SetApplicationStatus(status int32) {
	atomic.StoreInt32(&applicationStatus, status)
}

// ApplicationIsRunning 判断应用是否正在运行
func ApplicationIsRunning() bool {
	return atomic.LoadInt32(&applicationStatus) == ApplicationStatusRunning
}

// GetVersion 获取应用版本号
func GetVersion() string {
	return version
}

// SetVersion 设置应用版本号
func SetVersion(v string) {
	version = v
}

// GetBuildedAt 获取应用构建时间
func GetBuildedAt() string {
	return buildedAt
}

// SetBuildedAt 设置应用构建时间
func SetBuildedAt(v string) {
	buildedAt = v
}

// GetBasicConfig 获取基本配置信息
func GetBasicConfig() BasicConfig {
	prefix := "basic."
	basicConfig := BasicConfig{
		Name:         defaultViperX.GetString(prefix + "name"),
		RequestLimit: defaultViperX.GetUint(prefix + "requestLimit"),
		Listen:       defaultViperX.GetStringFromENV(prefix + "listen"),
	}
	validatePanic(&basicConfig)
	return basicConfig
}

// GetSessionConfig 获取session的配置
func GetSessionConfig() SessionConfig {
	prefix := "session."
	sessConfig := SessionConfig{
		TTL:        defaultViperX.GetDuration(prefix + "ttl"),
		Key:        defaultViperX.GetString(prefix + "key"),
		CookiePath: defaultViperX.GetString(prefix + "path"),
		Keys:       defaultViperX.GetStringSlice(prefix + "keys"),
		TrackKey:   defaultViperX.GetString(prefix + "trackKey"),
	}
	validatePanic(&sessConfig)
	return sessConfig
}

// GetRedisConfig 获取redis的配置
func GetRedisConfig() RedisConfig {
	prefix := "redis."
	uri := defaultViperX.GetStringFromENV(prefix + "uri")
	uriInfo, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	// 获取设置的db
	db := 0
	query := uriInfo.Query()
	dbValue := query.Get("db")
	if dbValue != "" {
		db, err = strconv.Atoi(dbValue)
		if err != nil {
			panic(err)
		}
	}
	// 获取密码
	password, _ := uriInfo.User.Password()

	// 获取slow设置的时间间隔
	slowValue := query.Get("slow")
	slow := 100 * time.Millisecond
	if slowValue != "" {
		slow, err = time.ParseDuration(slowValue)
		if err != nil {
			panic(err)
		}
	}

	// 获取最大处理数的配置
	maxProcessing := 1000
	maxValue := query.Get("maxProcessing")
	if maxValue != "" {
		maxProcessing, err = strconv.Atoi(maxValue)
		if err != nil {
			panic(err)
		}
	}

	redisConfig := RedisConfig{
		Addr:          uriInfo.Host,
		Password:      password,
		DB:            db,
		Slow:          slow,
		MaxProcessing: uint32(maxProcessing),
	}

	validatePanic(&redisConfig)
	return redisConfig
}

// GetMailConfig 获取邮件配置
func GetMailConfig() MailConfig {
	prefix := "mail."
	mailConfig := MailConfig{
		Host:     defaultViperX.GetString(prefix + "host"),
		Port:     defaultViperX.GetInt(prefix + "port"),
		User:     defaultViperX.GetStringFromENV(prefix + "user"),
		Password: defaultViperX.GetStringFromENV(prefix + "password"),
	}
	validatePanic(&mailConfig)
	return mailConfig
}

// GetInfluxdbConfig 获取influxdb配置
func GetInfluxdbConfig() InfluxdbConfig {
	prefix := "influxdb."
	influxdbConfig := InfluxdbConfig{
		URI:           defaultViperX.GetStringFromENV(prefix + "uri"),
		Bucket:        defaultViperX.GetString(prefix + "bucket"),
		Org:           defaultViperX.GetString(prefix + "org"),
		Token:         defaultViperX.GetStringFromENV(prefix + "token"),
		BatchSize:     defaultViperX.GetUint(prefix + "batchSize"),
		FlushInterval: defaultViperX.GetDuration(prefix + "flushInterval"),
		Disabled:      defaultViperX.GetBool(prefix + "disabled"),
	}
	validatePanic(&influxdbConfig)
	return influxdbConfig
}

// GetAlarmConfig get alarm config
func GetAlarmConfig() AlarmConfig {
	prefix := "alarm."
	alarmConfig := AlarmConfig{
		Receivers: defaultViperX.GetStringSlice(prefix + "receivers"),
	}
	validatePanic(&alarmConfig)
	return alarmConfig
}
