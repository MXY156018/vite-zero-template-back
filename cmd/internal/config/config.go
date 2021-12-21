package config

import "github.com/tal-tech/go-zero/rest"

type Server struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	JWT struct {
		SigningKey  string
		ExpiresTime int64
		BufferTime  int64
	}
	Zap    Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email  Email `mapstructure:"email" json:"email" yaml:"email"`
	Casbin struct {
		ModelPath string
	}
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
	Excel      Excel      `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer      Timer      `mapstructure:"timer" json:"timer" yaml:"timer"`
}
