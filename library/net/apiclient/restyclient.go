package apiclient

import (
	"time"

	xtime "focalboard-tool/library/time"

	mattermost "github.com/mattermost/mattermost-server/model"
)

const (
	DefaultTimeout = 10 * time.Second
)

type ClientConfig struct {
	Addr           string
	APIVersionPath string
	Token          string
	Timeout        xtime.Duration // HTTP客户端超时时间，支持从配置文件解析
}

type FRestClient struct {
	BaseURL string
	Token   string
	Timeout time.Duration // 保存超时配置
}

type MRestClient struct {
	Client  *mattermost.Client4
	Token   string
	Timeout time.Duration // 保存超时配置
}

func NewFRestClient(c *ClientConfig) *FRestClient {
	timeout := time.Duration(c.Timeout)
	if timeout == 0 {
		timeout = DefaultTimeout // 默认超时时间
	}

	return &FRestClient{
		BaseURL: c.Addr + c.APIVersionPath,
		Token:   c.Token,
		Timeout: timeout,
	}
}

func NewMRestClient(c *ClientConfig) *MRestClient {
	timeout := time.Duration(c.Timeout)
	if timeout == 0 {
		timeout = DefaultTimeout // 默认超时时间
	}

	client := mattermost.NewAPIv4Client(c.Addr + c.APIVersionPath)
	return &MRestClient{
		Client:  client,
		Token:   c.Token,
		Timeout: timeout,
	}
}
