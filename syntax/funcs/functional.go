package funcs

func Functional4() {
	println("Hello, I am Functional4")
}

func UseFunction4() {
	myFunc := Functional4
	myFunc()
}

// Functional5 函数作为返回值
func Functional5() func() string {
	return func() string {
		return "你好"
	}
}

// Functional6 匿名函数立即调用
func Functional6() {
	fn := func() string {
		return "hello"
	}()
	println(fn)
}
