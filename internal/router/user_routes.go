package router

import (
	"rich_go/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutes 设置用户相关路由
func SetupUserRoutes(v1 *gin.RouterGroup, handler *handlers.UserHandler) {
	users := v1.Group("/users")
	{
		users.GET("", handler.ListUsers)
		users.GET("/:id", handler.GetUser)
		users.POST("", handler.CreateUser)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}

