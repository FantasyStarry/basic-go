package main

import (
	"errors"
	"fmt"
)

// 实现切片的删除

// DeleteByIdx 删除指定位置的数据
func DeleteByIdx[T CompareElement](tSlice []T, idx int) ([]T, error) {
	if len(tSlice) == 0 || tSlice == nil {
		return nil, errors.New("切片不能为空")
	}
	if idx < 0 || idx > len(tSlice)-1 {
		return nil, errors.New("下标错误")
	}
	// 利用子切片
	res := make([]T, 0, len(tSlice)-1)
	frontSlice := tSlice[:idx]
	backendSlice := tSlice[idx+1:]
	res = append(frontSlice, backendSlice...)
	return res, nil
}

func DeleteByElement[T CompareElement](sourceSlice []T, element T) ([]T, error) {
	if len(sourceSlice) == 0 || sourceSlice == nil {
		return nil, errors.New("切片不能为空")
	}
	for idx, source := range sourceSlice {
		if source == element {
			return DeleteByIdx(sourceSlice, idx)
		}
	}
	return nil, errors.New("not found the element")
}

type CompareElement interface {
	int | int8 | int64 | int16 | int32 | string | bool
}

func main() {
	nameSlice := []string{
		"zhangsan", "wangwu", "laoliu",
	}
	//res, err := DeleteByIdx(nameSlice, 1)
	res, err := DeleteByElement(nameSlice, "zhangsan")
	if err != nil {
		fmt.Printf("error info is: %s", err)
	}
	fmt.Printf("deleted data is: %v", res)
}
