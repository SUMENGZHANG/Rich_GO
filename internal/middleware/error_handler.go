package middleware

import (
	"log"
	"rich_go/pkg/errors"
	"rich_go/pkg/response"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 统一错误处理中间件
// 注意：这个中间件主要用于捕获通过 c.Error() 设置的错误
// 实际错误处理主要在 Handler 层完成
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否有通过 c.Error() 设置的错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			// 记录错误日志
			log.Printf("Error: %v", err.Error())

			// 处理业务错误
			if be, ok := errors.AsBusinessError(err.Err); ok {
				response.Error(c, be.Code, be.Message)
				return
			}

			// 处理标准库错误
			if err.Err == errors.ErrRecordNotFound {
				response.NotFound(c, "资源不存在")
				return
			}

			// 默认错误处理
			response.InternalServerError(c, "内部服务器错误")
		}
	}
}

// Recovery 恢复中间件（增强版）
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log.Printf("Panic recovered: %v", recovered)
		response.InternalServerError(c, "服务器内部错误")
		c.Abort()
	})
}

