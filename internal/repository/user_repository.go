package repository

import (
	"context"
	"rich_go/internal/model"
)

// UserRepository 用户仓储接口
type UserRepository interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByID(ctx context.Context, id uint) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, id uint, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint) error
}

// userRepository 用户仓储实现（内存实现，后续可替换为数据库实现）
type userRepository struct {
	users  []*model.User
	nextID uint
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository() UserRepository {
	return &userRepository{
		users:  make([]*model.User, 0),
		nextID: 1,
	}
}

func (r *userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	// 返回副本，避免外部修改
	result := make([]*model.User, len(r.users))
	copy(result, r.users)
	return result, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			// 返回副本
			u := *user
			return &u, nil
		}
	}
	return nil, ErrNotFound
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	// 返回副本
	u := *user
	return &u, nil
}

func (r *userRepository) Update(ctx context.Context, id uint, user *model.User) (*model.User, error) {
	for i, u := range r.users {
		if u.ID == id {
			// 更新字段
			if user.Name != "" {
				r.users[i].Name = user.Name
			}
			if user.Email != "" {
				r.users[i].Email = user.Email
			}
			// 返回副本
			updated := *r.users[i]
			return &updated, nil
		}
	}
	return nil, ErrNotFound
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

