package service

import (
	"context"
	"focalboard-tool/pkg/errors"
)

// HandleNotFoundError 处理"资源不存在"错误的特殊情况
func (s *Service) HandleNotFoundError(err error, resourceType string, resourceID string) error {
	// 检查是否为"资源不存在"错误
	if errors.IsNotFoundError(err) {
		// 可以在这里添加特殊处理逻辑，比如查找备用资源、创建默认资源等
		// 例如记录到特殊日志、发送通知等

		// 这里我们只是将错误传递出去
		return err
	}

	// 不是"资源不存在"错误，直接传递
	return err
}

// HandleAuthError 处理认证/授权错误的特殊情况
func (s *Service) HandleAuthError(ctx context.Context, err error, userID string) error {
	// 检查是否为认证/授权错误
	if errors.IsAuthError(err) {
		// 可以在这里添加特殊处理逻辑
		// 例如记录到安全日志、通知管理员等

		// 这里我们只是将错误传递出去
		return err
	}

	// 不是认证/授权错误，直接传递
	return err
}
