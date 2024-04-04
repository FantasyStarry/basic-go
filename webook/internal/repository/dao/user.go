package dao

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

//数据库中的数据表结构

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	// 存毫秒数
	now := time.Now().UnixMilli()
	u.CreateTime = now
	u.UpdateTime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突，只有一个唯一索引
			return ErrUserDuplicateEmail
		}
	}
	return err
}

// User 直接对应数据库表结构
// 有些人叫做Entity 有些人叫做model 也有人叫做po
type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"` // 全部用户唯一
	Password string
	// 创建时间，毫秒数
	CreateTime int64
	// 更新时间，毫秒数
	UpdateTime int64
}
