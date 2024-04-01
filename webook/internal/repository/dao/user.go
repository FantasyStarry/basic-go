package dao

//数据库中的数据表结构

type User struct {
}

type Address struct {
	Id     int64
	UserId int64
}
