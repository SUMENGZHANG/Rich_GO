package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListUsers 获取用户列表
func ListUsers(c *gin.Context) {
	// TODO: 实现获取用户列表逻辑
	c.JSON(http.StatusOK, gin.H{
		"users": []gin.H{
			{"id": 1, "name": "用户1"},
			{"id": 2, "name": "用户2"},
		},
	})
}

// GetUser 获取单个用户
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: 从数据库获取用户信息
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": "用户" + id,
	})
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
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

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
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

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: 从数据库删除用户
	c.JSON(http.StatusOK, gin.H{
		"message": "用户删除成功",
		"id":      id,
	})
}

