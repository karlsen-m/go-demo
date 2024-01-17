package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("原始数组：")
	fmt.Println(arr)

	bubbleSort(arr)
	fmt.Println("冒泡排序后的数组：")
	fmt.Println(arr)

}
func bubbleSort(arr []int) {
	// 目标数组
	target := []int{22, 11}

	// 对目标数组进行排序
	sort.SliceStable(target, func(i, j int) bool {
		return target[i] < target[j]
	})

	// 对非目标数组元素进行倒序排序
	otherElements := []int{}
	for _, v := range arr {
		if !contains(target, v) {
			otherElements = append(otherElements, v)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(otherElements)))

	// 将排序后的目标数组和非目标数组合并
	result := append(target, otherElements...)

	copy(arr, result)
}

// 判断数组中是否包含某个元素
func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

type TopicType struct {
	GroupId   int    `json:"groupId"`
	Topic     string `json:"topic"`
	Questions []struct {
		Answer   string `json:"answer"`
		Sequence int    `json:"sequence"`
	} `json:"questions"`
}
