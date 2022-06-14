package test

import (
	"fmt"
	"sync"
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

func TestConcurrenceSlice1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)

	var mu sync.Mutex
	result := []int{}

	for i := 0; i < 5; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			result = append(result, i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	for _, v := range result {
		fmt.Println(v)
	}
}

func TestConcurrenceSlice2(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer close(ch1)
	defer close(ch2)

	result := []int{}

	go func(chan int) {
		ch1 <- 1
	}(ch1)

	go func(chan int) {
		ch2 <- 2
	}(ch2)

	result = append(result, <-ch1)
	result = append(result, <-ch2)

	for _, v := range result {
		fmt.Println(v)
	}
}

func TestConcurrenceSlice3(t *testing.T) {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	done := make(chan int, 5)
	defer close(ch1)
	defer close(ch2)
	defer close(done)

	result := []int{}

	go func(chan int) {
		ch1 <- 1
	}(ch1)

	go func(chan int) {
		ch2 <- 2
	}(ch2)

	count := 0

	for {
		select {
		case a := <-ch1:
			count++
			result = append(result, a)
			fmt.Printf("ch1 count==%d\n", count)
			fmt.Println(result)
			if count == 2 {
				done <- 1
			}
		case b := <-ch2:
			count++
			result = append(result, b)
			fmt.Printf("ch2 count==%d\n", count)
			fmt.Println(result)
			if count == 2 {
				done <- 1
			}
		case c := <-done:
			if c == 1 {
				fmt.Printf("done==%d\n", count)
				goto out
			}
		}
	}

out:
	for _, v := range result {
		fmt.Println(v)
	}
}

func TestConcurrenceSlice4(t *testing.T) {
	ch := make(chan []int, 5)
	defer close(ch)

	result := []int{}

	go func() {
		array := []int{1, 2, 3, 4, 5}
		ch <- array
	}()

	v := <-ch
	for _, va := range v {
		result = append(result, va)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

func TestConcurrenceSlice5(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan []int, 5)
	result := []int{}

	go func() {
		arr := []int{1, 2, 3, 4, 5}
		ch <- arr
		wg.Done()
	}()

	go func() {
		arr := []int{6, 7, 8, 9, 10}
		ch <- arr
		wg.Done()
	}()

	wg.Wait()

	// Go提供了range关键字，将其使用在channel上时，会自动等待channel的动作一直到channel被关闭
	// 所以使用range遍历channel之前，必须确保channel已关闭，否则会死锁
	close(ch)

	for va := range ch {
		for _, v := range va {
			result = append(result, v)
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

// 查找数组中最大的前k个数
// 最简单的方案是排序后取出前k个数
// 因为选择排序每次都能挑出最大值
// 所以可以改造选择排序算法实现，外层循环只需循环k次，时间复杂度为O(k*n)
func searchTopK(array []int, k int) []int {
	for i := 0; i < k; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] > array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array[:k]
}

func TestTopK(t *testing.T) {
	arr := []int{7, 4, 5, 3, 1, 8, 6, 6, 9}
	result := searchTopK(arr, 5)
	fmt.Println(result)
}

// 去除有序数组重复元素：使用快慢指针
func rmDuplicateArray(arr []int) {
	slow, fast := 0, 1
	for fast < len(arr) {
		if arr[fast] != arr[slow] {
			slow++
			// 维护 arr[0..slow] ⽆重复,神来之笔,将重复元素替换为后边的非重复元素
			arr[slow] = arr[fast]
		}
		fast++
	}
	// ⻓度为索引 + 1
	fmt.Println(arr[:slow+1])
}

func TestRmDuplicates(t *testing.T) {
	arr := []int{1,1,1,2,2,2,3,3,4,5,6,6,7}
	rmDuplicateArray(arr)
}

// leetcode-55:给定一个非负整数数组，数组中每个元素代表你在该位置可以跳跃的最大距离，
// 从数组第一个位置开始，判断能否到达最后一个位置
// 应用贪心算法：每⼀步都做出⼀个局部最优的选择，最终的结果就是全局最优。
func judgeJump2last(arr []int) bool {
	n, far := len(arr), 0
	for i:=0;i<n-1;i++ {
		// 不断计算能跳到的最远距离
		far = Max(far, i+arr[i])
		// 可能碰到了 0，卡住跳不动了
		if far <= i {
			return false
		}
	}
	return far > n-1
}

func TestJump(t *testing.T) {
	arr := []int{3,2,1,0,4,5}
	res := judgeJump2last(arr)
	t.Log(res)
}


// leetcode-27:给定一个num数组和一个数字，在仅使用O(1)额外空间的情况下删除数组中该数字,返回删除后数组长度
// 如：nums = [0,1,2,2,3,0,4,2], val = 2,删除后数组变为[0,1,3,0,4],返回5
func removeElements(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	// 记录相等数字的下标位置
	count := 0
	for i:=0;i<len(nums);i++ {
		if nums[i] != val {
			if i != count {
				// 交换位置，相等元素后移，不等元素前移
				nums[i], nums[count] = nums[count], nums[i]
			}
			count++
		}
	}
	fmt.Println(nums[:count])
	return count
}

func TestRemoveElements(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	count := removeElements(nums, val)
	t.Log(count)
}


// 给定⼀个不重复的排序数组和⼀个⽬标值，在数组中找到⽬标值，并返回其索引。
// 如果⽬标值不存在于数组中，返回它将会被按顺序插⼊的位置。
// 二分查找变种题
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1 //右移一位相当于除2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}