package server

import (
	"net/http"

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
	router.GET("/health", healthCheck)

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 用户相关路由
		users := v1.Group("/users")
		{
			users.GET("", listUsers)
			users.GET("/:id", getUser)
			users.POST("", createUser)
			users.PUT("/:id", updateUser)
			users.DELETE("/:id", deleteUser)
		}
	}
}

// Start 启动 HTTP 服务器
func (s *HTTPServer) Start() error {
	return s.router.Run(":" + s.port)
}

// healthCheck 健康检查接口
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "服务运行正常",
	})
}

// listUsers 获取用户列表
func listUsers(c *gin.Context) {
	// TODO: 实现获取用户列表逻辑
	c.JSON(http.StatusOK, gin.H{
		"users": []gin.H{
			{"id": 1, "name": "用户1"},
			{"id": 2, "name": "用户2"},
		},
	})
}

// getUser 获取单个用户
func getUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: 从数据库获取用户信息
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": "用户" + id,
	})
}

// createUser 创建用户
func createUser(c *gin.Context) {
	var user struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: 保存用户到数据库
	c.JSON(http.StatusCreated, gin.H{
		"message": "用户创建成功",
		"user":    user,
	})
}

// updateUser 更新用户
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email" binding:"omitempty,email"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: 更新用户信息
	c.JSON(http.StatusOK, gin.H{
		"message": "用户更新成功",
		"id":      id,
		"user":    user,
	})
}

// deleteUser 删除用户
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: 从数据库删除用户
	c.JSON(http.StatusOK, gin.H{
		"message": "用户删除成功",
		"id":      id,
	})
}

