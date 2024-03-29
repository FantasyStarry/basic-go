package main

import "fmt"

func Byte() {
	var a byte = 'a'
	println(a)
	fmt.Printf("%c", a)

	var str string = "this is string"
	var bs []byte = []byte(str)
	println(str, bs)
}
