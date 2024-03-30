package generics

func Sum[T Number](args ...T) T {
	var res T
	for _, arg := range args {
		res = res + arg
	}
	return res
}

type Number interface {
	int | int64
}
