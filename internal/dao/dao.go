package dao

import (
	"focalboard-tool/internal/conf"
	xhttp "focalboard-tool/library/net/apiclient"
	"time"

	fbClient "github.com/mattermost/focalboard/server/client"
)

// Dao struct
type Dao struct {
	c            *conf.Config
	restyFClient *xhttp.FRestClient
	restyMClient *xhttp.MRestClient
}

// New init
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c:            c,
		restyFClient: xhttp.NewFRestClient(c.HttpClient.FocalboardClient),
		restyMClient: xhttp.NewMRestClient(c.HttpClient.MattermostClient),
	}
	return
}

// createFocalboardClient 创建并配置FocalBoard客户端
// 这是一个通用方法，用于统一创建客户端并设置通用配置
func (d *Dao) createFocalboardClient(token string) *fbClient.Client {
	rclientUrl := d.restyFClient.BaseURL
	rclient := fbClient.NewClient(rclientUrl, token)
	rclient.HTTPClient.Timeout = d.restyFClient.Timeout
	return rclient
}

// GetFocalboardClientTimeout 获取FocalBoard客户端配置的超时时间
// 这个方法可以用于调试和验证配置
func (d *Dao) GetFocalboardClientTimeout() time.Duration {
	return d.restyFClient.Timeout
}

// GetFocalboardClientBaseURL 获取FocalBoard客户端配置的基础URL
// 这个方法可以用于调试和验证配置
func (d *Dao) GetFocalboardClientBaseURL() string {
	return d.restyFClient.BaseURL
}
