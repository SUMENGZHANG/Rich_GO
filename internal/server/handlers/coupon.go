package handlers

import (
	"rich_go/internal/service"
	"rich_go/pkg/errors"
	"rich_go/pkg/response"

	"github.com/gin-gonic/gin"
)

// CouponHandler 优惠券处理器
type CouponHandler struct {
	couponService service.CouponService
}

// NewCouponHandler 创建优惠券处理器实例
func NewCouponHandler(couponService service.CouponService) *CouponHandler {
	return &CouponHandler{
		couponService: couponService,
	}
}

// ListCoupons 获取优惠券列表
func (h *CouponHandler) ListCoupons(c *gin.Context) {
	coupons, err := h.couponService.ListCoupons(c.Request.Context())
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			response.Error(c, be.Code, be.Message)
			return
		}
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, gin.H{"coupons": coupons})
}

// GetCoupon 获取单个优惠券
func (h *CouponHandler) GetCoupon(c *gin.Context) {
	id := c.Param("id")
	coupon, err := h.couponService.GetCoupon(c.Request.Context(), id)
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			if be.Code == errors.CodeCouponNotFound || be.Code == errors.CodeInvalidCouponID {
				response.NotFound(c, be.Message)
			} else {
				response.Error(c, be.Code, be.Message)
			}
			return
		}
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, coupon)
}

// CreateCoupon 创建优惠券
func (h *CouponHandler) CreateCoupon(c *gin.Context) {
	var req struct {
		Name         string  `json:"name" binding:"required"`
		Description  string  `json:"description"`
		DiscountType string  `json:"discountType" binding:"required,oneof=fixed percent"`
		DiscountValue float64 `json:"discountValue" binding:"required,gt=0"`
		MinAmount    float64 `json:"minAmount" binding:"gte=0"`
		Status       string  `json:"status" binding:"omitempty,oneof=active inactive"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	createReq := &service.CreateCouponRequest{
		Name:         req.Name,
		Description:  req.Description,
		DiscountType: req.DiscountType,
		DiscountValue: req.DiscountValue,
		MinAmount:    req.MinAmount,
		Status:       req.Status,
	}

	coupon, err := h.couponService.CreateCoupon(c.Request.Context(), createReq)
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			response.Error(c, be.Code, be.Message)
			return
		}
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "优惠券创建成功", coupon)
}

// UpdateCoupon 更新优惠券
func (h *CouponHandler) UpdateCoupon(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		DiscountType string  `json:"discountType" binding:"omitempty,oneof=fixed percent"`
		DiscountValue float64 `json:"discountValue" binding:"omitempty,gt=0"`
		MinAmount    float64 `json:"minAmount" binding:"omitempty,gte=0"`
		Status       string  `json:"status" binding:"omitempty,oneof=active inactive"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	updateReq := &service.UpdateCouponRequest{
		Name:         req.Name,
		Description:  req.Description,
		DiscountType: req.DiscountType,
		DiscountValue: req.DiscountValue,
		MinAmount:    req.MinAmount,
		Status:       req.Status,
	}

	coupon, err := h.couponService.UpdateCoupon(c.Request.Context(), id, updateReq)
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			if be.Code == errors.CodeCouponNotFound || be.Code == errors.CodeInvalidCouponID {
				response.NotFound(c, be.Message)
			} else {
				response.Error(c, be.Code, be.Message)
			}
			return
		}
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "优惠券更新成功", coupon)
}

// DeleteCoupon 删除优惠券
func (h *CouponHandler) DeleteCoupon(c *gin.Context) {
	id := c.Param("id")
	err := h.couponService.DeleteCoupon(c.Request.Context(), id)
	if err != nil {
		if be, ok := errors.AsBusinessError(err); ok {
			if be.Code == errors.CodeCouponNotFound || be.Code == errors.CodeInvalidCouponID {
				response.NotFound(c, be.Message)
			} else {
				response.Error(c, be.Code, be.Message)
			}
			return
		}
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "优惠券删除成功", gin.H{"id": id})
}

