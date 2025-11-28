package router

import (
	"rich_go/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有路由
func SetupRoutes(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
	couponHandler *handlers.CouponHandler,
) {
	// 健康检查接口
	router.GET("/health", handlers.HealthCheck)

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		SetupUserRoutes(v1, userHandler)
		SetupCouponRoutes(v1, couponHandler)
	}
}

