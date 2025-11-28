package handlers

import (
	"rich_go/pkg/response"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查接口
func HealthCheck(c *gin.Context) {
	response.SuccessWithMessage(c, "服务运行正常", gin.H{
		"status": "ok",
	})
}

