package main

import "math"

func main() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)
	// 取余
	println(a % b)
	// 只有同类型的数据才能进行加减乘除
	var c float64 = 123.123
	println(a + int(c))

	// 使用math包
	math.Abs(123)
	String()

	Byte()
}
