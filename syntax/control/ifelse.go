package control

func IfOnly(age int) {
	if age > 18 {
		println("已经成年")
	}
}

func IfElse(age int) {
	if age > 18 {
		println("成年的孩子")
	} else {
		println("还是一个孩子")
	}
}
