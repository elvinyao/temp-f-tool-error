package service

import (
	"context"
	"focalboard-tool/library/log"
	"focalboard-tool/pkg/utils"

	"go.uber.org/zap"
)

// LogServiceParams 记录服务方法参数
func LogServiceParams(ctx context.Context, methodName string, params map[string]interface{}) {
	// 从上下文中获取请求ID
	requestID := "unknown"
	if id, exists := ctx.Value("X-Request-ID").(string); exists {
		requestID = id
	}

	// 使用通用的敏感信息处理工具
	sanitizedParams := utils.SanitizeParams(params)

	// 记录日志
	log.Info("服务方法调用",
		zap.String("request_id", requestID),
		zap.String("method", methodName),
		zap.Any("params", sanitizedParams),
	)
}
