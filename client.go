package chengquan

import (
	"errors"
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string
	AppID  string
	AppKey string
	AesKey string
	AesIv  string
}

// Client 实例
type Client struct {
	config struct {
		apiURL  string
		appID   string
		appKey  string
		aesKey  string
		aesIv   string
		version string
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}
	if config.ApiURL == "" {
		return nil, errors.New("需要配置ApiURL")
	}

	c.config.apiURL = config.ApiURL
	c.config.appID = config.AppID
	c.config.appKey = config.AppKey
	c.config.aesKey = config.AesKey
	c.config.aesIv = config.AesKey
	c.config.version = "1.0.0"

	return c, nil
}
