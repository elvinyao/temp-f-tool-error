package http

import "github.com/gin-gonic/gin"

func (r apiroutes) addFocalboardApi(rg *gin.RouterGroup) {
	focalboardApi := rg.Group("/focalboard")

	// Get a single board by its ID.
	focalboardApi.GET("/boards/single", SingleBoard)

}
