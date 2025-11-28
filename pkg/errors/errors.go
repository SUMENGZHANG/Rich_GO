package errors

import (
	"errors"
	"fmt"
)

// 业务错误码定义
const (
	CodeSuccess = 0

	// 通用错误码 1000-1999
	CodeInvalidParam  = 1001
	CodeNotFound      = 1002
	CodeInternalError = 1003

	// 用户相关错误码 2000-2999
	CodeUserNotFound     = 2001
	CodeUserAlreadyExists = 2002
	CodeInvalidUserID    = 2003

	// 优惠券相关错误码 3000-3999
	CodeCouponNotFound      = 3001
	CodeCouponAlreadyExists = 3002
	CodeInvalidCouponID     = 3003
	CodeInvalidDiscountType = 3004
)

// BusinessError 业务错误
type BusinessError struct {
	Code    int
	Message string
}

func (e *BusinessError) Error() string {
	return e.Message
}

// NewBusinessError 创建业务错误
func NewBusinessError(code int, message string) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: message,
	}
}

// NewBusinessErrorf 创建格式化业务错误
func NewBusinessErrorf(code int, format string, args ...interface{}) *BusinessError {
	return &BusinessError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// 预定义错误
var (
	ErrInvalidParam  = NewBusinessError(CodeInvalidParam, "参数错误")
	ErrNotFound      = NewBusinessError(CodeNotFound, "资源不存在")
	ErrInternalError = NewBusinessError(CodeInternalError, "内部服务器错误")

	ErrUserNotFound     = NewBusinessError(CodeUserNotFound, "用户不存在")
	ErrUserAlreadyExists = NewBusinessError(CodeUserAlreadyExists, "用户已存在")
	ErrInvalidUserID    = NewBusinessError(CodeInvalidUserID, "无效的用户ID")

	ErrCouponNotFound      = NewBusinessError(CodeCouponNotFound, "优惠券不存在")
	ErrCouponAlreadyExists = NewBusinessError(CodeCouponAlreadyExists, "优惠券已存在")
	ErrInvalidCouponID     = NewBusinessError(CodeInvalidCouponID, "无效的优惠券ID")
	ErrInvalidDiscountType = NewBusinessError(CodeInvalidDiscountType, "无效的折扣类型")
)

// IsBusinessError 判断是否为业务错误
func IsBusinessError(err error) bool {
	_, ok := err.(*BusinessError)
	return ok
}

// AsBusinessError 转换为业务错误
func AsBusinessError(err error) (*BusinessError, bool) {
	be, ok := err.(*BusinessError)
	return be, ok
}

// WrapError 包装错误
func WrapError(err error, code int, message string) error {
	if err == nil {
		return nil
	}
	return NewBusinessError(code, fmt.Sprintf("%s: %v", message, err))
}

// 标准库错误兼容
var (
	ErrRecordNotFound = errors.New("record not found")
)

