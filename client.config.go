package chengquan

import (
	"go.dtapp.net/gorequest"
)

func (c *Client) SetApiURL(v string) *Client {
	c.config.apiURL = v
	return c
}

func (c *Client) SetAppID(v string) *Client {
	c.config.appID = v
	return c
}

func (c *Client) SetAppKey(v string) *Client {
	c.config.appKey = v
	return c
}

func (c *Client) SetAesKey(v string) *Client {
	c.config.aesKey = v
	return c
}

func (c *Client) SetAesIv(v string) *Client {
	c.config.aesIv = v
	return c
}

// SetClientIP 配置
func (c *Client) SetClientIP(clientIP string) *Client {
	c.clientIP = clientIP
	if c.httpClient != nil {
		c.httpClient.SetClientIP(clientIP)
	}
	return c
}

// SetLogFun 设置日志记录函数
func (c *Client) SetLogFun(logFun gorequest.LogFunc) {
	if c.httpClient != nil {
		c.httpClient.SetLogFunc(logFun)
	}
}
