package domain

import "time"

//操作实际业务上的数据

// User 领域对象，是DDD中的Entity
type User struct {
	Id         int64
	Email      string
	Password   string
	CreateTime time.Time
}

//type Address struct {
//}
