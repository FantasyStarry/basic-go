package main

type List interface {
	Add(idx int, val int) error
	Append(val int)
	Delete(idx int) error
}
