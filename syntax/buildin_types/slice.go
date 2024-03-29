package main

import "fmt"

func Slice1() {
	s1 := []int{1, 2, 3}
	fmt.Printf("s1 = %v, len = %d, cap = %d \n", s1, len(s1), cap(s1))

	// 长度3，容量4
	s2 := make([]int, 3, 4)
	fmt.Printf("s2 = %v, len = %d, cap = %d \n", s2, len(s2), cap(s2))

	// 长度和容量都是4
	s3 := make([]int, 4)
	fmt.Printf("s3 = %v, len = %d, cap = %d \n", s3, len(s3), cap(s3))

	s4 := make([]int, 0, 4)
	s4 = append(s4, 1)
	fmt.Printf("s4 = %v, len = %d, cap = %d \n", s4, len(s4), cap(s4))
}

func SubSlice() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[1:3]
	fmt.Printf("s2 = %v, len = %d, cap = %d \n", s2, len(s2), cap(s2))
	s2[1] = 10
	s2 = append(s2, 110)
	fmt.Printf("s2 = %v, len = %d, cap = %d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1 = %v, len = %d, cap = %d \n", s1, len(s1), cap(s1))
}

func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s2 = %v, len = %d, cap = %d \n", s2, len(s2), cap(s2))
	s1[0] = 99
	fmt.Printf("s1 = %v, len = %d, cap = %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len = %d, cap = %d \n", s2, len(s2), cap(s2))
	s2[0] = 99
	fmt.Printf("s1 = %v, len = %d, cap = %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len = %d, cap = %d \n", s2, len(s2), cap(s2))
	// 如果发生了扩容，就会更变底层数组，就不会产生切片共享
	s2 = append(s2, 1999)
	fmt.Printf("s1 = %v, len = %d, cap = %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len = %d, cap = %d \n", s2, len(s2), cap(s2))
}
