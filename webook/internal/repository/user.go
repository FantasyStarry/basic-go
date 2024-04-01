package repository

// UserRepository 定义了用户仓库的结构体
// 这个结构体用于存储和操作用户相关数据
type UserRepository struct {
	// 以下是用户仓库内部的字段定义
	// ...
}

func (r *UserRepository) FindById(userId int64) {
	// 先从cache中查找
	// 再从dao里面找
	// 找到了再写会cache

}
