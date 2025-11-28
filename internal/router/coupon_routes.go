package router

import (
	"rich_go/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

// SetupCouponRoutes 设置优惠券相关路由
func SetupCouponRoutes(v1 *gin.RouterGroup, handler *handlers.CouponHandler) {
	coupons := v1.Group("/coupons")
	{
		coupons.GET("", handler.ListCoupons)
		coupons.GET("/:id", handler.GetCoupon)
		coupons.POST("", handler.CreateCoupon)
		coupons.PUT("/:id", handler.UpdateCoupon)
		coupons.DELETE("/:id", handler.DeleteCoupon)
	}
}

