package repository

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository/dao"
	"context"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail

// UserRepository 定义了用户仓库的结构体
// 这个结构体用于存储和操作用户相关数据
type UserRepository struct {
	// 以下是用户仓库内部的字段定义
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	// 在这里操作缓存
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindById(userId int64) {
	// 先从cache中查找
	// 再从dao里面找
	// 找到了再写会cache

}
