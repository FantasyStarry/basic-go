package service

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

var (
	ErrUserDuplicateEmail    = repository.ErrUserDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("邮箱或密码错误")
	ErrUserNotFound          = repository.ErrUserNotFound
)

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 考虑加密存放的问题
	// 然后就是存起来
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, u domain.User) (domain.User, error) {
	// 先找用户
	uRepo, err := svc.repo.FindByEmail(ctx, u)
	if errors.Is(err, repository.ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(uRepo.Password), []byte(u.Password))
	if err != nil {
		return uRepo, ErrInvalidUserOrPassword
	}
	return uRepo, nil
}

func (svc *UserService) Profile(ctx *gin.Context, userId int64) (domain.User, error) {
	userRepo, err := svc.repo.FindById(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	return userRepo, nil
}
