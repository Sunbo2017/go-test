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
// 可以参考快排思想实现
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
			// 维护 arr[0..slow] ⽆重复,神来之笔,将重复元素替换为后边的非重复元素，即重复元素后移
			arr[slow] = arr[fast]
		}
		fast++
	}
	// ⻓度为索引 + 1
	fmt.Println(arr[:slow+1])
}

func TestRmDuplicates(t *testing.T) {
	arr := []int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 6, 6, 7}
	rmDuplicateArray(arr)
}

// leetcode-27:给定一个num数组和一个数字，在仅使用O(1)额外空间的情况下删除数组中该数字,返回删除后数组长度
// 如：nums = [0,1,2,2,3,0,4,2], val = 2,删除后数组变为[0,1,3,0,4],返回5
func removeElements(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	// 记录相等数字的下标位置
	count := 0
	for i := 0; i < len(nums); i++ {
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
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	count := removeElements(nums, val)
	t.Log(count)
}

// leetcode-55:给定一个非负整数数组，数组中每个元素代表你在该位置可以跳跃的最大距离，
// 从数组第一个位置开始，判断能否到达最后一个位置
// 应用贪心算法：每⼀步都做出⼀个局部最优的选择，最终的结果就是全局最优。
func judgeJump2last(arr []int) bool {
	n, far := len(arr), 0
	for i := 0; i < n-1; i++ {
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
	arr := []int{3, 2, 1, 0, 4, 5}
	res := judgeJump2last(arr)
	t.Log(res)
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
			//遍历到了最后一位 或 mid后一位正好大于等于目标值
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}

// 快速模幂运算
// 要求你的算法返回幂运算  a^b  的计算结果与 1337 取模（mod，也就是余数）后的结果。
// 就是你先得计算幂  a^b  ，但是这个  b  会⾮常⼤，所以  b 是⽤数组的形式表⽰的。
// 公式：(a * b) % k = (a % k)(b % k) % k
var base = 1337

// 计算 a 的 k 次⽅然后与 base 求模的结果
func mypow(a, k int) int {
	// 对因⼦求模
	a %= base
	res := 1
	for i := 0; i < k; i++ {
		// 这⾥有乘法，是潜在的溢出点
		res *= a
		// 对乘法结果求模
		res %= base
	}
	return res
}

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]

	part1 := mypow(a, last)
	part2 := mypow(superPow(a, b), 10)

	return (part1 * part2) % base
}

// 在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，
// 每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
// [
// [1,2,8,9],
// [2,4,9,12],
// [4,7,10,13],
// [6,8,11,15]
// ]
// 给定 target = 7，返回 true。
// 给定 target = 3，返回 false。
// 暴力解法，双重循环，直接查找
// 采用分治法，二分查找变种：
// step 1：首先获取矩阵的两个边长，判断特殊情况。
// step 2：首先以左下角为起点，若它小于目标元素，则往右移动去找大的，若他大于目标元素，则往上移动去找小的。
// step 3：若是移动到了矩阵边界也没找到，说明矩阵中不存在目标值。
func find2Array(arr [][]int, target int) bool {
	a := len(arr)
	b := len(arr[0])
	if a == 0 || b == 0 {
		return false
	}
	for i, j := a-1, 0; i >= 0 && j < b; {
		//元素较大，往上走
		if arr[i][j] > target {
			i--
		} else if arr[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

// 给定一个长度为n的数组，返回其中任何一个峰值的索引
// 峰值元素是指其值严格大于左右相邻值的元素
// 二分查找
func findHighest(arr []int) int {
	low, high := 0, len(arr)-1
	for low < high {
		mid := low + (high-low)>>1
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		} else if arr[mid] > arr[mid+1] && arr[mid] < arr[mid-1] {
			//右边是往下，不一定有坡峰,所以往左走
			high = mid
		} else {
			//右边是往上，一定能找到波峰
			low = mid + 1
		}
	}
	return high
}

// 买卖股票最佳时机,只可以买入卖出一次
// 贪心：循环记录每次收益，结果取最大值
func getMaxProfit1(prices []int) int {
	min := 10000
	max := 0
	for _, v := range prices {
		min = Min(min, v)
		max = Max(max, v-min)
	}
	return max
}

//买卖股票最佳时机，可以买入卖出多次
//贪心：只要收益为正就累加收益
func getMaxProfit2(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

//买卖股票最佳时机，最多可买入卖出两次
func getMaxProfit3(prices []int) int {
	// dp定义
	// buy1 代表第1次买入的股票收益
	// sell1代表第1次卖出的股票收益
	// buy2 代表第2次买入的股票收益
	// sell2代表第2次卖出的股票收益

	// 基本状态
	// buy1=-price, sell1=0
	// buy2=-price, sell2=0
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		// 转移方程
		buy1 = Max(buy1, -price)
		sell1 = Max(sell1, buy1+price)
		buy2 = Max(buy2, sell1-price)
		sell2 = Max(sell2, buy2+price)
	}
	return sell2
}

func josephus(n int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	m := 1
	for len(nums) > 1 {
		for i := 0; i < len(nums); i++ {
			if m > 3 {
				m = 1
			}
			if m%3 == 0 {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
			m++
		}
	}

	return nums[0]
}

func TestJosephus(t *testing.T) {
	r := josephus(5)
	t.Log(r)
}
