package test

import (
	"fmt"
	"testing"
)

func TestMaxArrayChild(t *testing.T) {
	arr := []int{15, 7, 4, 8, 12, 120, 100, 200, 300}
	result := maxArrayChild(arr)
	fmt.Println("result============")
	fmt.Println(result)
}

// 最大连续递增子序列
func maxArrayChild(arr []int) []int {
	arr1 := []int{}
	// arrmap := map[int][]int{}
	max := 0
	result := []int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] > arr[i] {
			arr1 = append(arr1, arr[i+1])
		} else {
			if len(arr1) > 1 {
				// arrmap[len(arr1)] = arr1
				if len(arr1) > max {
					result = arr1
					max = len(arr1)
				}
				fmt.Println("test=====")
				fmt.Println(arr1)
			}
			// 创建一个新数组并添加第一个元素
			arr1 = []int{arr[i+1]}
		}
	}
	if len(arr1) > 1 {
		// arrmap[len(arr1)] = arr1
		if len(arr1) > max {
			result = arr1
			max = len(arr1)
		}
		fmt.Println("test1=====")
		fmt.Println(arr1)
	}
	// fmt.Println(arrmap)
	// for k := range arrmap {
	// 	if k > max {
	// 		max = k
	// 	}
	// }
	// fmt.Println(result)
	return result
}
