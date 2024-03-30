package component

import "fmt"

type Inner struct {
}

func (i Inner) DoSomethings() {
	fmt.Printf("这是 inner")
}

type Outer struct {
	Inner
}

type OuterV1 struct {
	Inner
}

func (v OuterV1) DoSomethings() {
	fmt.Printf("这个 outer v1")
}

type OuterPtr struct {
	*Inner
}

func UseInner() {
	var o Outer
	o.DoSomethings()

	var op OuterPtr
	op.DoSomethings()

	o1 := Outer{
		Inner: Inner{},
	}
	o1.DoSomethings()

	op1 := OuterPtr{
		Inner: &Inner{},
	}
	op1.DoSomethings()

	ov1 := OuterV1{
		Inner: Inner{},
	}
	//如果出现相同的先找自己的，有点类似重写，但不是重写
	ov1.DoSomethings()       // 这是outer v1
	ov1.Inner.DoSomethings() // 这是 inner
}
