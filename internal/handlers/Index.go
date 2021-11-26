package handlers

import (
	"github.com/gin-gonic/gin"
	"sheetServerApi/internal/middlewares/response"
)

// 测试接口联通
func Index(c *gin.Context) {
	response.ResponseSuccess(c,200,"ok","接口正常")
}

