package service

import (
	"context"
	"rich_go/internal/model"
	"rich_go/internal/repository"
	"rich_go/pkg/errors"
	"strconv"
)

// CouponService 优惠券服务接口
type CouponService interface {
	ListCoupons(ctx context.Context) ([]*model.Coupon, error)
	GetCoupon(ctx context.Context, idStr string) (*model.Coupon, error)
	CreateCoupon(ctx context.Context, req *CreateCouponRequest) (*model.Coupon, error)
	UpdateCoupon(ctx context.Context, idStr string, req *UpdateCouponRequest) (*model.Coupon, error)
	DeleteCoupon(ctx context.Context, idStr string) error
}

// CreateCouponRequest 创建优惠券请求
type CreateCouponRequest struct {
	Name         string
	Description  string
	DiscountType string
	DiscountValue float64
	MinAmount    float64
	Status       string
}

// UpdateCouponRequest 更新优惠券请求
type UpdateCouponRequest struct {
	Name         string
	Description  string
	DiscountType string
	DiscountValue float64
	MinAmount    float64
	Status       string
}

// couponService 优惠券服务实现
type couponService struct {
	couponRepo repository.CouponRepository
}

// NewCouponService 创建优惠券服务实例
func NewCouponService(couponRepo repository.CouponRepository) CouponService {
	return &couponService{
		couponRepo: couponRepo,
	}
}

func (s *couponService) ListCoupons(ctx context.Context) ([]*model.Coupon, error) {
	return s.couponRepo.FindAll(ctx)
}

func (s *couponService) GetCoupon(ctx context.Context, idStr string) (*model.Coupon, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, errors.ErrInvalidCouponID
	}
	coupon, err := s.couponRepo.FindByID(ctx, uint(id))
	if err == repository.ErrNotFound {
		return nil, errors.ErrCouponNotFound
	}
	return coupon, err
}

func (s *couponService) CreateCoupon(ctx context.Context, req *CreateCouponRequest) (*model.Coupon, error) {
	// 业务逻辑验证
	if req.Name == "" {
		return nil, errors.NewBusinessError(errors.CodeInvalidParam, "优惠券名称不能为空")
	}
	if req.DiscountType != "fixed" && req.DiscountType != "percent" {
		return nil, errors.ErrInvalidDiscountType
	}
	if req.DiscountValue <= 0 {
		return nil, errors.NewBusinessError(errors.CodeInvalidParam, "折扣值必须大于 0")
	}
	if req.MinAmount < 0 {
		return nil, errors.NewBusinessError(errors.CodeInvalidParam, "最低使用金额不能小于 0")
	}

	// 设置默认状态
	status := req.Status
	if status == "" {
		status = "active"
	}

	coupon := &model.Coupon{
		Name:         req.Name,
		Description:  req.Description,
		DiscountType: req.DiscountType,
		DiscountValue: req.DiscountValue,
		MinAmount:    req.MinAmount,
		Status:       status,
	}

	return s.couponRepo.Create(ctx, coupon)
}

func (s *couponService) UpdateCoupon(ctx context.Context, idStr string, req *UpdateCouponRequest) (*model.Coupon, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, errors.ErrInvalidCouponID
	}

	// 业务逻辑验证
	if req.DiscountType != "" && req.DiscountType != "fixed" && req.DiscountType != "percent" {
		return nil, errors.ErrInvalidDiscountType
	}
	if req.DiscountValue < 0 {
		return nil, errors.NewBusinessError(errors.CodeInvalidParam, "折扣值不能小于 0")
	}
	if req.MinAmount < 0 {
		return nil, errors.NewBusinessError(errors.CodeInvalidParam, "最低使用金额不能小于 0")
	}

	coupon := &model.Coupon{
		Name:         req.Name,
		Description:  req.Description,
		DiscountType: req.DiscountType,
		DiscountValue: req.DiscountValue,
		MinAmount:    req.MinAmount,
		Status:       req.Status,
	}

	result, err := s.couponRepo.Update(ctx, uint(id), coupon)
	if err == repository.ErrNotFound {
		return nil, errors.ErrCouponNotFound
	}
	return result, err
}

func (s *couponService) DeleteCoupon(ctx context.Context, idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return errors.ErrInvalidCouponID
	}
	err = s.couponRepo.Delete(ctx, uint(id))
	if err == repository.ErrNotFound {
		return errors.ErrCouponNotFound
	}
	return err
}

