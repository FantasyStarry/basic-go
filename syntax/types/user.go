package main

import "fmt"

type User struct {
	Name string
	age  int8
}

func (u User) ChangeName(name string) {
	// 值传递，复制了一个u，修改的是复制的数据
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	u.age = int8(age)
}

func NewUser() {
	//初始化结构体
	u := User{}
	fmt.Printf("u %+v \n", u)

	up := &User{}
	fmt.Printf("up %+v \n", up)

	up2 := new(User)
	fmt.Printf("up2 %+v \n", up2)

	u4 := User{Name: "Tom", age: 123} //优先选择
	u5 := User{"Tome", 23}

	u4.Name = "Jerry"
	u5.age = 18
}

func ChangeUser() {
	u1 := User{
		Name: "Tom",
		age:  20,
	}
	u1.ChangeName("Jerry")
	u1.ChangeAge(18)
}
