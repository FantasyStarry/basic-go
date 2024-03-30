package generics

// List T是参数，名字叫做T，约束是any等于没有约束
type List[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UseList() {
	var l List[int]
	l.Append(10)
}
