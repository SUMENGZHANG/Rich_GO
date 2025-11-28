package server

import (
	"rich_go/internal/middleware"
	"rich_go/internal/repository"
	"rich_go/internal/router"
	"rich_go/internal/server/handlers"
	"rich_go/internal/service"

	"github.com/gin-gonic/gin"
)

// HTTPServer HTTP 服务器结构
type HTTPServer struct {
	router *gin.Engine
	port   string
}

// NewHTTPServer 创建新的 HTTP 服务器实例（使用依赖注入）
func NewHTTPServer(port string) *HTTPServer {
	// 设置 Gin 模式
	// gin.SetMode(gin.ReleaseMode) // 生产环境使用
	gin.SetMode(gin.DebugMode) // 开发环境使用

	engine := gin.Default()

	// 添加全局中间件
	setupMiddleware(engine)

	// 初始化 Repository 层
	userRepo := repository.NewUserRepository()
	couponRepo := repository.NewCouponRepository()

	// 初始化 Service 层
	userService := service.NewUserService(userRepo)
	couponService := service.NewCouponService(couponRepo)

	// 初始化 Handler 层
	userHandler := handlers.NewUserHandler(userService)
	couponHandler := handlers.NewCouponHandler(couponService)

	// 注册路由
	router.SetupRoutes(engine, userHandler, couponHandler)

	return &HTTPServer{
		router: engine,
		port:   port,
	}
}

// setupMiddleware 设置中间件
func setupMiddleware(router *gin.Engine) {
	// 使用自定义恢复中间件
	router.Use(middleware.Recovery())
	
	// 使用自定义日志中间件（可选，Gin 默认也有）
	// router.Use(middleware.Logger())
	
	// 错误处理中间件
	router.Use(middleware.ErrorHandler())
	
	// 可以在这里添加其他中间件
	// router.Use(middleware.Auth())
	// router.Use(middleware.CORS())
}

// Start 启动 HTTP 服务器
func (s *HTTPServer) Start() error {
	return s.router.Run(":" + s.port)
}

