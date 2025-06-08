package dao

import (
	"focalboard-tool/pkg/errors"

	fbClient "github.com/mattermost/focalboard/server/client"
)

// handleConfigFocalboardError 使用配置处理Focalboard API错误
func (d *Dao) handleConfigFocalboardError(operation string, resourceID string, respBody *fbClient.Response, result interface{}) error {
	if respBody.Error == nil && respBody.StatusCode == 200 {
		return nil
	}

	if respBody.StatusCode == 404 {
		return errors.ConfigResourceNotFound(resourceID, respBody.Error)
	} else if respBody.StatusCode == 401 || respBody.StatusCode == 403 {
		return errors.ConfigUnauthorized(respBody.Error)
	} else {
		return errors.ConfigFocalboardAPIError(operation, respBody.Error)
	}
}
