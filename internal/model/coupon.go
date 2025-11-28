package model

// Coupon 优惠券模型
type Coupon struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	DiscountType string  `json:"discountType"` // fixed: 固定金额, percent: 百分比
	DiscountValue float64 `json:"discountValue"`
	MinAmount    float64 `json:"minAmount"`
	Status       string  `json:"status"` // active: 启用, inactive: 禁用
}

