package repository

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository/dao"
	"context"
	"time"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

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

func (r *UserRepository) FindByEmail(ctx context.Context, u domain.User) (domain.User, error) {
	email := u.Email
	ud, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       ud.Id,
		Email:    ud.Email,
		Password: ud.Password,
	}, nil
}

func (r *UserRepository) FindById(ctx context.Context, userId int64) (domain.User, error) {
	// 先从cache中查找
	// 再从dao里面找
	// 找到了再写会cache
	ud, err := r.dao.FindById(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	birthday := time.Unix(ud.Birthday, 0)
	return domain.User{
		Id:       ud.Id,
		Email:    ud.Email,
		Nickname: ud.Nickname,
		Birthday: birthday,
		AutoMe:   ud.AutoMe,
	}, nil
}

func (r *UserRepository) EditUserInfo(ctx context.Context, user domain.User) error {
	ud := dao.User{
		Id:       user.Id,
		Nickname: user.Nickname,
		Birthday: user.Birthday.Unix(),
		AutoMe:   user.AutoMe,
	}
	return r.dao.UpdateUserInfo(ctx, ud)
}
