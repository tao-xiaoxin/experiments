package main

import (
	"errors"
	"fmt"
)

// Delete 按照索引删除单个指定元素
func Delete[T any](slice []T, index int) ([]T, T, error) {
	length := len(slice)
	if (index < 0) || (index >= length) {
		var zero T
		return nil, zero, errors.New("index out of range")
	}
	fmt.Println("Element:", slice[index])
	ele := slice[index]
	return append(slice[:index], slice[index+1:]...), ele, nil
}

// Remove 按照切片中的某个值删除
func Remove[T comparable](slice []T, value T) ([]T, error) {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...), nil
		}
	}
	return nil, errors.New("value not found")
}

func main() {
	//// 示例使用
	//intList := []int{1, 2, 3, 4, 5}
	//indexToDelete := 1
	//
	//updatedList, deletedElement, err := Delete(intList, indexToDelete)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//
	//fmt.Println("Updated List:", updatedList)
	//fmt.Println("Deleted Element:", deletedElement)

	intSlice := []int{1, 2, 3, 4, 5}
	valueToRemove := 3

	updatedSlice, err := Remove(intSlice, valueToRemove)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Updated Slice:", updatedSlice)

}
