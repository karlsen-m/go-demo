package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []float64{100, 110.1, 110.2, 201, 205, 101, 110.1, 111.2, 111.2, 202, 204}
	sort.Float64s(slice1) // 排序切片

	target := 111.22
	index := binarySearchEQ(slice1, target)

	fmt.Printf("小于 %v 的数据索引： %v\n", target, index)
}
func binarySearchEQ(arr []float64, target float64) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid // 返回目标元素的索引
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // 如果未找到目标元素，则返回 -1
}

func binarySearchLt(arr []float64, target float64) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func binarySearchLET(arr []float64, target float64) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
