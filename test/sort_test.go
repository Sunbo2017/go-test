package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 冒泡排序
func Bubble(arr []int) {
	size := len(arr)
	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped != true {
			break
		}
	}
}

// 选择排序
func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
				// fmt.Println(arr)
			}
		}
		// fmt.Printf("loop %v \n", i)
	}
}

// 插入排序
func InsertSort(arr []int) {
	for i := 1; i <= len(arr)-1; i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

// 快速排序
func QuickSort(arr []int, l, r int) {
	if l < r {
		pivot := arr[r]
		i := l - 1
		for j := l; j < r; j++ {
			if arr[j] <= pivot {
				i++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		i++
		arr[r], arr[i] = arr[i], arr[r]
		QuickSort(arr, l, i-1)
		QuickSort(arr, i+1, r)
	}
}

// 合并
func Merge(arr []int, l, mid, r int) {
	// 分别复制左右子数组
	n1, n2 := mid-l+1, r-mid
	left, right := make([]int, n1), make([]int, n2)
	copy(left, arr[l:mid+1])
	copy(right, arr[mid+1:r+1])
	i, j := 0, 0
	k := l
	for ; i < n1 && j < n2; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	for ; i < n1; i++ {
		arr[k] = left[i]
		k++
	}
	for ; j < n2; j++ {
		arr[k] = right[j]
		k++
	}
}

//归并排序
func MergeSort(arr []int, l, r int) {
	if l < r {
		mid := (l + r - 1) / 2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, r)
		Merge(arr, l, mid, r)
	}
}

//堆调整
func adjustHeap(arr []int, i, size int) {
	if i <= (size-2)/2 {
		//左右子节点
		l, r := 2*i+1, 2*i+2
		m := i
		if l < size && arr[l] > arr[m] {
			m = l
		}
		if r < size && arr[r] > arr[m] {
			m = r
		}
		if m != i {
			arr[m], arr[i] = arr[i], arr[m]
			adjust_heap(arr, m, size)
		}
	}
}

//建堆
func buildHeap(arr []int) {
	size := len(arr)
	//从最后一个子节点开始向前调整
	for i := (size - 2) / 2; i >= 0; i-- {
		adjust_heap(arr, i, size)
	}
}

// 堆排序
func HeapSort(arr []int) {
	size := len(arr)
	buildHeap(arr)
	for i := size - 1; i > 0; i-- {
		//顶部arr[0]为当前最大值,调整到数组末尾
		arr[0], arr[i] = arr[i], arr[0]
		adjustHeap(arr, 0, i)
	}
}

func TestMySort(t *testing.T) {
	//数量级，控制参与排序的数字总量
	num := 10000
	array := []int{}
	// array := []int{5, 6, 3, 2, 1, 0, 9, 7, 8, 10, 20, 50, 21, 16, 12, 18, 23, 30, 40, 32}
	// array := []int{5, 6, 3, 2, 1, 0, 9, 7, 8}
	// 随机种子，不加的话每次产生的随机数相同
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		array = append(array, rand.Intn(num))
	}

	start := time.Now().UnixNano()
	fmt.Printf("start time: %v \n", start)
	// fmt.Println(array)

	// SelectSort(array) //数量级1000：3000300 ns 2998000 3007200 数量级10000：236098400
	// Bubble(array) //数量级1000：3000400 ns 3000500 3001600 3000700 2009100 1994100 数量级10000：234680800 231802700
	// InsertSort(array) //数量级1000：2000600 ns 1999600 2001000 2013600 数量级10000：148584200
	// MergeSort(array, 0, len(array)-1) //数量级10000：3008000
	// HeapSort(array) //数量级10000：2000200
	QuickSort(array, 0, len(array)-1) //数量级10000：1000300

	// fmt.Println(array)
	end := time.Now().UnixNano()
	fmt.Printf("end time: %v \n", end)
	fmt.Printf("time cost: %v", end-start)
}
