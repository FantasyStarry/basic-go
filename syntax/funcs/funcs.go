package main

// Func1 可以没有参数
func Func1() {

}

// Func2 可以有一个参数
func Func2(name string) {

}

// Func3 可以有多个参数
func Func3(name string, password string) {

}

func Func4(name string) string {
	return "可以是一个返回值"
}

func Func5() (name string, msg string) {
	return "Spongzi", "可以是多个返回值"
}

func Func6() (name string, msg string) {
	// 返回值带名字可以直接return
	return
}
