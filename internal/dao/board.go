package dao

import (
	"context"

	fbModel "github.com/mattermost/focalboard/server/model"
)

// ConfigSingleBoard 使用配置的错误处理获取单个看板
func (d *Dao) SingleBoard(c context.Context, boardId string, token string) (*fbModel.Board, error) {
	rclient := d.createFocalboardClient(token)
	boardres, respBody := rclient.GetBoard(boardId, rclient.Token)

	// 使用基于配置的错误处理函数
	if err := d.handleConfigFocalboardError("GetBoard", boardId, respBody, boardres); err != nil {
		return nil, err
	}

	return boardres, nil
}
