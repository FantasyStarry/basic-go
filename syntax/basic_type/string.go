package main

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	println("Hello String")
	println(`你好
			我可以换行
			再来一行
			`)

	// 字符串的拼接用+且只能和string进行拼接不能直接转义
	println("abc" + "cba")
	//println("abc" + string(123))
	println(fmt.Sprintf("abc%d", 123))

	// len计算的是字节数
	println(len("abc"))
	println(len("你好")) // 输出6
	println(utf8.RuneCountInString("你好"))

	//strings包

}
