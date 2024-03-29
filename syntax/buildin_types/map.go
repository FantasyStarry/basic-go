package main

import "fmt"

func Map() {
	m1 := map[string]string{
		"key": "value",
	}
	// 修改
	m1["key"] = "123"

	// 容量
	m2 := make(map[string]string, 12)
	m2["张三"] = "男"
	val, ok := m2["王五"]
	if ok {
		fmt.Printf("Get the value is %s", val)
	}
	val = m2["张三"]
	fmt.Printf("张三对应的值是:%s", val)

	delete(m2, "张三")
}
