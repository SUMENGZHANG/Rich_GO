package repository

import (
	"context"
	"rich_go/internal/model"
)

// CouponRepository 优惠券仓储接口
type CouponRepository interface {
	FindAll(ctx context.Context) ([]*model.Coupon, error)
	FindByID(ctx context.Context, id uint) (*model.Coupon, error)
	Create(ctx context.Context, coupon *model.Coupon) (*model.Coupon, error)
	Update(ctx context.Context, id uint, coupon *model.Coupon) (*model.Coupon, error)
	Delete(ctx context.Context, id uint) error
}

// couponRepository 优惠券仓储实现（内存实现，后续可替换为数据库实现）
type couponRepository struct {
	coupons []*model.Coupon
	nextID  uint
}

// NewCouponRepository 创建优惠券仓储实例
func NewCouponRepository() CouponRepository {
	return &couponRepository{
		coupons: make([]*model.Coupon, 0),
		nextID:  1,
	}
}

func (r *couponRepository) FindAll(ctx context.Context) ([]*model.Coupon, error) {
	result := make([]*model.Coupon, len(r.coupons))
	copy(result, r.coupons)
	return result, nil
}

func (r *couponRepository) FindByID(ctx context.Context, id uint) (*model.Coupon, error) {
	for _, coupon := range r.coupons {
		if coupon.ID == id {
			c := *coupon
			return &c, nil
		}
	}
	return nil, ErrNotFound
}

func (r *couponRepository) Create(ctx context.Context, coupon *model.Coupon) (*model.Coupon, error) {
	coupon.ID = r.nextID
	r.nextID++
	r.coupons = append(r.coupons, coupon)
	c := *coupon
	return &c, nil
}

func (r *couponRepository) Update(ctx context.Context, id uint, coupon *model.Coupon) (*model.Coupon, error) {
	for i, c := range r.coupons {
		if c.ID == id {
			if coupon.Name != "" {
				r.coupons[i].Name = coupon.Name
			}
			if coupon.Description != "" {
				r.coupons[i].Description = coupon.Description
			}
			if coupon.DiscountType != "" {
				r.coupons[i].DiscountType = coupon.DiscountType
			}
			if coupon.DiscountValue > 0 {
				r.coupons[i].DiscountValue = coupon.DiscountValue
			}
			if coupon.MinAmount >= 0 {
				r.coupons[i].MinAmount = coupon.MinAmount
			}
			if coupon.Status != "" {
				r.coupons[i].Status = coupon.Status
			}
			updated := *r.coupons[i]
			return &updated, nil
		}
	}
	return nil, ErrNotFound
}

func (r *couponRepository) Delete(ctx context.Context, id uint) error {
	for i, coupon := range r.coupons {
		if coupon.ID == id {
			r.coupons = append(r.coupons[:i], r.coupons[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

