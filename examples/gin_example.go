package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin REST API 示例
func main() {
	// 创建 Gin 路由引擎
	// gin.Default() 包含 Logger 和 Recovery 中间件
	// gin.New() 创建不带中间件的引擎
	r := gin.Default()

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "服务运行正常",
		})
	})

	// GET 接口示例
	r.GET("/api/v1/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": "用户" + id,
		})
	})

	// POST 接口示例
	r.POST("/api/v1/users", func(c *gin.Context) {
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

		c.JSON(http.StatusCreated, gin.H{
			"message": "用户创建成功",
			"user":   user,
		})
	})

	// 路由组示例
	v1 := r.Group("/api/v1")
	{
		v1.GET("/products", getProducts)
		v1.GET("/products/:id", getProduct)
		v1.POST("/products", createProduct)
	}

	// 启动服务器，默认监听 0.0.0.0:8080
	r.Run(":8080")
}

// 获取产品列表
func getProducts(c *gin.Context) {
	products := []gin.H{
		{"id": 1, "name": "产品1", "price": 99.99},
		{"id": 2, "name": "产品2", "price": 199.99},
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

// 获取单个产品
func getProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"name":  "产品" + id,
		"price": 99.99,
	})
}

// 创建产品
func createProduct(c *gin.Context) {
	var product struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "产品创建成功",
		"product": product,
	})
}

