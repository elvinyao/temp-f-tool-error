package service

import (
	"context"
	"focalboard-tool/internal/apimodel"
)

// FocalboardService 定义Focalboard服务接口
type FocalboardService interface {
	// SingleBoard 获取单个看板
	SingleBoard(c context.Context, token interface{}, board interface{}) (*apimodel.Board, error)

	// 下面是未来可能添加的方法
	// GetCard(c context.Context, token interface{}, cardID interface{}) (*apimodel.Card, error)
	// UpdateCard(c context.Context, token interface{}, cardID interface{}, card *apimodel.Card) error
	// CreateCard(c context.Context, token interface{}, boardID interface{}, card *apimodel.Card) (*apimodel.Card, error)
	// DeleteCard(c context.Context, token interface{}, cardID interface{}) error
	// GetBoardCards(c context.Context, token interface{}, boardID interface{}) ([]*apimodel.Card, error)
}
