package funcs

// Closure 闭包
func Closure(name string) func() string {
	return func() string {
		return name + "hello"
	}
}

func Closure1() func() string {
	name := "大米"
	return func() string {
		return name + "hello"
	}
}
