package main

//按ctrl + i 可以直接实现接口

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Add(idx int, val int) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(val int) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(idx int) error {
	//TODO implement me
	panic("implement me")
}

type Node struct {
}
