package server

import (
	"rich_go/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

// HTTPServer HTTP 服务器结构
type HTTPServer struct {
	router *gin.Engine
	port   string
}

// NewHTTPServer 创建新的 HTTP 服务器实例
func NewHTTPServer(port string) *HTTPServer {
	// 设置 Gin 模式
	// gin.SetMode(gin.ReleaseMode) // 生产环境使用
	gin.SetMode(gin.DebugMode) // 开发环境使用

	router := gin.Default()

	// 添加全局中间件
	setupMiddleware(router)

	// 注册路由
	setupRoutes(router)

	return &HTTPServer{
		router: router,
		port:   port,
	}
}

// setupMiddleware 设置中间件
func setupMiddleware(router *gin.Engine) {
	// Gin 默认中间件已包含 Logger 和 Recovery
	// 可以在这里添加自定义中间件
	// router.Use(customMiddleware())
}

// setupRoutes 设置路由
func setupRoutes(router *gin.Engine) {
	// 健康检查接口
	router.GET("/health", handlers.HealthCheck)

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 用户相关路由
		setupUserRoutes(v1)
		
		// 优惠券相关路由
		setupCouponRoutes(v1)
	}
}

// setupUserRoutes 设置用户相关路由
func setupUserRoutes(v1 *gin.RouterGroup) {
	users := v1.Group("/users")
	{
		users.GET("", handlers.ListUsers)
		users.GET("/:id", handlers.GetUser)
		users.POST("", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
}

// setupCouponRoutes 设置优惠券相关路由
func setupCouponRoutes(v1 *gin.RouterGroup) {
	coupons := v1.Group("/coupons")
	{
		coupons.GET("", handlers.ListCoupons)
		coupons.GET("/:id", handlers.GetCoupon)
		coupons.POST("", handlers.CreateCoupon)
		coupons.PUT("/:id", handlers.UpdateCoupon)
		coupons.DELETE("/:id", handlers.DeleteCoupon)
	}
}

// Start 启动 HTTP 服务器
func (s *HTTPServer) Start() error {
	return s.router.Run(":" + s.port)
}

