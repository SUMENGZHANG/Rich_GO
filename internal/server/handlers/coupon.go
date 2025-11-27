package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListCoupons 获取优惠券列表
func ListCoupons(c *gin.Context) {
	// TODO: 实现获取优惠券列表逻辑
	c.JSON(http.StatusOK, gin.H{
		"coupons": []gin.H{
			{
				"id":           1,
				"name":         "新用户专享",
				"description":  "新用户首次购买立减10元",
				"discountType": "fixed",
				"discountValue": 10,
				"minAmount":    50,
				"status":       "active",
			},
			{
				"id":           2,
				"name":         "满减优惠",
				"description":  "满100减20",
				"discountType": "fixed",
				"discountValue": 20,
				"minAmount":    100,
				"status":       "active",
			},
		},
	})
}

// GetCoupon 获取单个优惠券
func GetCoupon(c *gin.Context) {
	id := c.Param("id")
	// TODO: 从数据库获取优惠券信息
	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"name":         "优惠券" + id,
		"description":  "这是优惠券" + id + "的描述",
		"discountType": "fixed",
		"discountValue": 10,
		"minAmount":    50,
		"status":       "active",
	})
}

// CreateCoupon 创建优惠券
func CreateCoupon(c *gin.Context) {
	var coupon struct {
		Name         string  `json:"name" binding:"required"`
		Description  string  `json:"description"`
		DiscountType string  `json:"discountType" binding:"required,oneof=fixed percent"`
		DiscountValue float64 `json:"discountValue" binding:"required,gt=0"`
		MinAmount    float64 `json:"minAmount" binding:"gte=0"`
		Status       string  `json:"status" binding:"omitempty,oneof=active inactive"`
	}

	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 设置默认状态
	if coupon.Status == "" {
		coupon.Status = "active"
	}

	// TODO: 保存优惠券到数据库
	c.JSON(http.StatusCreated, gin.H{
		"message": "优惠券创建成功",
		"coupon":  coupon,
	})
}

// UpdateCoupon 更新优惠券
func UpdateCoupon(c *gin.Context) {
	id := c.Param("id")
	var coupon struct {
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		DiscountType string  `json:"discountType" binding:"omitempty,oneof=fixed percent"`
		DiscountValue float64 `json:"discountValue" binding:"omitempty,gt=0"`
		MinAmount    float64 `json:"minAmount" binding:"omitempty,gte=0"`
		Status       string  `json:"status" binding:"omitempty,oneof=active inactive"`
	}

	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: 更新优惠券信息
	c.JSON(http.StatusOK, gin.H{
		"message": "优惠券更新成功",
		"id":      id,
		"coupon":  coupon,
	})
}

// DeleteCoupon 删除优惠券
func DeleteCoupon(c *gin.Context) {
	id := c.Param("id")
	// TODO: 从数据库删除优惠券
	c.JSON(http.StatusOK, gin.H{
		"message": "优惠券删除成功",
		"id":      id,
	})
}

