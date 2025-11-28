package service

import (
	"context"
	"rich_go/internal/model"
	"rich_go/internal/repository"
	"rich_go/pkg/errors"
	"strconv"
)

// UserService 用户服务接口
type UserService interface {
	ListUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, idStr string) (*model.User, error)
	CreateUser(ctx context.Context, name, email string) (*model.User, error)
	UpdateUser(ctx context.Context, idStr string, name, email string) (*model.User, error)
	DeleteUser(ctx context.Context, idStr string) error
}

// userService 用户服务实现
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.userRepo.FindAll(ctx)
}

func (s *userService) GetUser(ctx context.Context, idStr string) (*model.User, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, errors.ErrInvalidUserID
	}
	user, err := s.userRepo.FindByID(ctx, uint(id))
	if err == repository.ErrNotFound {
		return nil, errors.ErrUserNotFound
	}
	return user, err
}

func (s *userService) CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	// 业务逻辑验证
	if name == "" {
		return nil, errors.NewBusinessError(errors.CodeInvalidParam, "用户名不能为空")
	}
	if email == "" {
		return nil, errors.NewBusinessError(errors.CodeInvalidParam, "邮箱不能为空")
	}

	user := &model.User{
		Name:  name,
		Email: email,
	}

	return s.userRepo.Create(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, idStr string, name, email string) (*model.User, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, errors.ErrInvalidUserID
	}

	user := &model.User{
		Name:  name,
		Email: email,
	}

	result, err := s.userRepo.Update(ctx, uint(id), user)
	if err == repository.ErrNotFound {
		return nil, errors.ErrUserNotFound
	}
	return result, err
}

func (s *userService) DeleteUser(ctx context.Context, idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return errors.ErrInvalidUserID
	}
	err = s.userRepo.Delete(ctx, uint(id))
	if err == repository.ErrNotFound {
		return errors.ErrUserNotFound
	}
	return err
}

