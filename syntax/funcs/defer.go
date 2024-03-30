package main

import "fmt"

// Defer 后进先出
func Defer() {
	defer func() {
		println("第一天defer")
	}()

	defer func() {
		println("第二个defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

type MyStruct struct {
	name string
}

func DeferClosureStruct() *MyStruct {
	var a = &MyStruct{
		name: "张三",
	}
	defer func() {
		a.name = "王五"
	}()
	return a
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("before: i 的地址是: %p, 值是：%d \n", &i, i) // 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
		// defer先进后出
		defer func() {
			fmt.Printf("i 的地址是: %p, 值是：%d \n", &i, i) // 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("i 的地址是: %p, 值是：%d \n", &val, val) // 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j) // 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
		}()
	}
}
