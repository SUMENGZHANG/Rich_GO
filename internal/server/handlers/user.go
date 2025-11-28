package handlers

import (
	"rich_go/internal/service"
	"rich_go/pkg/errors"
	"rich_go/pkg/response"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// ListUsers 获取用户列表
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.userService.ListUsers(c.Request.Context())
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			response.Error(c, be.Code, be.Message)
			return
		}
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, gin.H{"users": users})
}

// GetUser 获取单个用户
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUser(c.Request.Context(), id)
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			if be.Code == errors.CodeUserNotFound || be.Code == errors.CodeInvalidUserID {
				response.NotFound(c, be.Message)
			} else {
				response.Error(c, be.Code, be.Message)
			}
			return
		}
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, user)
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			response.Error(c, be.Code, be.Message)
			return
		}
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "用户创建成功", user)
}

// UpdateUser 更新用户
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email" binding:"omitempty,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), id, req.Name, req.Email)
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			if be.Code == errors.CodeUserNotFound || be.Code == errors.CodeInvalidUserID {
				response.NotFound(c, be.Message)
			} else {
				response.Error(c, be.Code, be.Message)
			}
			return
		}
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "用户更新成功", user)
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.userService.DeleteUser(c.Request.Context(), id)
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			if be.Code == errors.CodeUserNotFound || be.Code == errors.CodeInvalidUserID {
				response.NotFound(c, be.Message)
			} else {
				response.Error(c, be.Code, be.Message)
			}
			return
		}
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "用户删除成功", gin.H{"id": id})
}

