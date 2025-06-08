package http

import (
	"focalboard-tool/internal/appconst"
	"focalboard-tool/internal/middleware"
	"focalboard-tool/pkg/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

//	@Summary	Get a single board by its ID.
//	@Produce	json
//	@Param		boardId	query		string	true	"Board Id"
//	@Param		token	query		string	true	"User token"
//	@Success	200		{object}	middleware.SuccessResponse
//	@Failure	400		{object}	middleware.ErrorResponse
//	@Failure	500		{object}	middleware.ErrorResponse
//	@Router		/api/v1/focalboard/boards/single [get]
func SingleBoard(c *gin.Context) {
	boardId := c.Query("boardId")
	token := c.Query("token")

	// 改进的参数验证
	if err := validateSingleBoardParams(boardId, token); err != nil {
		c.Error(err)
		return
	}

	// 调用服务
	sboard, err := srv.SingleBoard(c, token, boardId)
	if err != nil {
		c.Error(err)
		return
	}

	// 成功响应
	middleware.RespondSuccess(c, sboard)
}

// validateSingleBoardParams 验证SingleBoard接口的参数
func validateSingleBoardParams(boardId, token string) error {
	// 验证token
	if token == "" {
		return errors.ConfigMissingParam("token")
	}

	// 验证token格式（基本检查）
	if len(strings.TrimSpace(token)) < appconst.ParamLengthLimit {
		return errors.ConfigInvalidParam("token", "token长度不足", nil)
	}

	// 验证boardId
	if boardId == "" {
		return errors.ConfigMissingParam("boardId")
	}

	// 验证boardId格式（基本检查）
	if len(strings.TrimSpace(boardId)) < appconst.ParamLengthLimit {
		return errors.ConfigInvalidParam("boardId", "boardId格式无效", nil)
	}

	// 检查是否包含非法字符
	if strings.ContainsAny(boardId, "<>\"'&") {
		return errors.ConfigInvalidParam("boardId", "包含非法字符", nil)
	}

	return nil
}
