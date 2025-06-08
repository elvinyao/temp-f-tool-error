package service

import (
	"context"

	"focalboard-tool/internal/apimodel"
)

// ConfigSingleBoard 使用配置化错误处理的获取看板方法
func (s *Service) SingleBoard(c context.Context, token interface{}, board interface{}) (*apimodel.Board, error) {
	// 使用辅助函数记录服务方法参数
	LogServiceParams(c, "ConfigSingleBoard", map[string]interface{}{
		"token": token,
		"board": board,
	})

	// 参数验证 - 使用配置化的验证函数
	boardStr, err := ConfigValidateStringParam("board", board)
	if err != nil {
		return nil, err
	}

	tokenStr, err := ConfigValidateStringParam("token", token)
	if err != nil {
		return nil, err
	}

	// 调用使用配置错误的DAO层方法
	oneboard, err := s.dao.SingleBoard(c, boardStr, tokenStr)
	if err != nil {
		// 错误已经在DAO层处理，直接传递
		return nil, err
	}

	// 转换为API响应模型
	singleBoard := &apimodel.Board{
		ID:             oneboard.ID,
		TeamID:         oneboard.TeamID,
		Title:          oneboard.Title,
		CardProperties: oneboard.CardProperties,
	}

	return singleBoard, nil
}
