package test

import (
	"fmt"
	"testing"
	"time"
)

// 冒泡排序：依次比较相邻两元素，每轮循环挑出最大值置于数组末尾
// 时间复杂度：O(N^2)
func Bubble(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 选择排序：第一个元素依次与每个元素比较，每轮循环挑出最小值置于数组开头
// 时间复杂度：O(N^2)
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

// 插入排序：认为第一个元素为有序数列，其他元素依次与之前的元素比较，大于插入右侧，小于插入左侧
// 时间复杂度：O(N^2)
func InsertSort(arr []int) {
	for i := 1; i <= len(arr)-1; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				// 前边元素一直有序，所以当出现不满足的元素时可以直接跳出内层循环
				break
			}
		}
	}
}

//计数排序
//当输入的元素是 n 个 0 到 k 之间的整数时，它的运行时间是 O(n + k)。计数排序不是比较排序，排序的速度快于任何比较排序算法。
//算法的步骤如下：
//（1）找出待排序的数组中最大和最小的元素
//（2）统计数组中每个值为i的元素出现的次数，存入数组C的第i项
//（3）对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
//（4）反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1
func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组

	sortedIndex := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}

// 快速排序：选中一个基准值，使得其左侧所有元素小于它，右侧大于它，从基准值位置将数组分为两部分，递归排序
// 时间复杂度：O(N*logN)
func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi)
	QuickSort(a, lo, p-1)
	QuickSort(a, p+1, hi)
}

func partitionSort(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] <= pivot {
			//i为慢指针，记录小于基准值的元素位置索引
			i++
			//小数左移
			a[j], a[i] = a[i], a[j]
		}
	}
	//确定基准值的位置并置换
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

// 快速排序：这种写法更易理解
// 时间复杂度：O(N*logN)
func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		// pivot := arr[(start+end)/2]
		pivot := arr[start]
		for i <= j {
			for arr[i] < pivot {
				i++
			}
			for arr[j] > pivot {
				j--
			}
			// 大数右移，小数左移
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		//经过最后一次循环计算，j缩小为中轴前一位，i增大为中轴后一位
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

// 归并排序，算法是采用分治法（Divide and Conquer）的一个非常典型的应用，且各层分治递归可以同时进行。
// 当我们要排序这样一个数组的时候，归并排序法首先将这个数组分成一半。
// 然后想办法把左边的数组给排序，右边的数组给排序，之后呢再将它们归并起来。
// 对左边的数组和右边的数组进行排序的时候，再分别将左边的数组和右边的数组分成一半，然后对每一个部分先排序，再归并。
// 分到一定细度的时候，每一个部分都只有一个元素，此时不用排序，对他们进行一次简单的归并就好了。
// 归并到上一个层级之后继续归并，直到归并到最高的层级，排序结束。
// 时间复杂度：O(N*logN)
func MergeSort(arr []int, l, r int) {
	if l < r {
		mid := (l + r - 1) / 2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, r)
		Merge(arr, l, mid, r)
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
			adjustHeap(arr, m, size)
		}
	}
}

//建堆
func buildHeap(arr []int) {
	size := len(arr)
	//从最后一个子节点开始向前调整
	for i := (size - 2) / 2; i >= 0; i-- {
		adjustHeap(arr, i, size)
	}
}

// 堆排序
// 堆是具有以下性质的完全二叉树：
// 每个结点的值都大于或等于其左右孩子结点的值，称为大顶堆；
// 或者每个结点的值都小于或等于其左右孩子结点的值，称为小顶堆。
// 时间复杂度：O(N*logN)
func HeapSort(arr []int) {
	size := len(arr)
	buildHeap(arr)
	for i := size - 1; i > 0; i-- {
		//顶部arr[0]为当前最大值,调整到数组末尾
		arr[0], arr[i] = arr[i], arr[0]
		adjustHeap(arr, 0, i)
	}
}

//数量级小与1000时，各种排序均可在1ns之内完成，无法统计时间；平方复杂度下，10000数量级插入排序最快
//但是在万级元素下，快排和插入排序并没有拉开太大差距，十万数量级下快排优势明显
func TestMySort(t *testing.T) {
	//数量级，控制参与排序的数字总量
	//num := 100000
	//array := []int{}
	//// array := []int{5, 6, 3, 2, 1, 0, 9, 7, 8}
	//// 随机种子，不加的话每次产生的随机数相同
	//rand.Seed(time.Now().Unix())
	//for i := 0; i < num; i++ {
	//	r := rand.Intn(num)
	//	//fmt.Println(r)
	//	array = append(array, r)
	//}
	//
	//fmt.Println(array)

	array1 := []int{5, 6, 3, 2, 1, 0, 9, 7, 8, 10, 20}

	start := time.Now().UnixNano()
	fmt.Printf("start time: %v \n", start)

	//SelectSort(array) //数量级1000：3000300 ns 2998000 3007200 数量级10000：123961100
	//Bubble(array) //数量级1000：3000400 ns 3000500 3001600 3000700 2009100 1994100 数量级10000：75448300
	//InsertSort(array1) //数量级1000：2000600 ns 1999600 2001000 2013600 数量级10000：40950500; 100000: 4338811300
	// MergeSort(array, 0, len(array)-1) //数量级10000：3008000
	// HeapSort(array) //数量级10000：2000200
	//QuickSort(array, 0, len(array)-1) //数量级10000：1000300
	QuickSort(array1, 0, len(array1)-1) //数量级10000：1000300; 100000: 5340800

	fmt.Println(array1)
	end := time.Now().UnixNano()
	fmt.Printf("end time: %v \n", end)
	fmt.Printf("time cost: %v ns\n", end-start)

	//fmt.Println(array)
}

// 二分查找
func binarySearch(sortedArray []int, target int) int {
	low := 0
	high := len(sortedArray) - 1
	for low <= high {
		mid := low + (high-low)/2
		midValue := sortedArray[mid]
		if midValue < target {
			low = mid + 1
		} else if midValue > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
//给定一个整数数组 a ,同时给定它的大小n和要找的 k
// step 1：进行一次快排，大元素在左，小元素在右，得到的中轴p点。
// step 2：如果 p - low + 1 = k ，那么p点就是第K大。
// step 3：如果 p - low + 1 > k，则第k大的元素在左半段，更新high = p - 1，执行step 1。
// step 4：如果 p - low + 1 < k，则第k大的元素在右半段，更新low = p + 1, 且 k = k - (p - low + 1)，排除掉前面部分更大的元素，再执行step 1.
func findTopK(arr []int, n, k int) int {
	return quickSort2(arr, 0, n-1, k)
}

func TestFindTopK(t *testing.T) {
	arr := []int{5, 3, 6, 9, 2, 1, 4}
	res := findTopK(arr, len(arr), 5)
	t.Log(res)
}

func quickSort2(arr []int, low, high, k int) int {
	p := partitionSortDesc(arr, low, high)
	if p-low+1 == k {
		return arr[p]
	} else if p-low+1 > k {
		//递归左边
		return quickSort2(arr, low, p-1, k)
	} else {
		//递归右边
		return quickSort2(arr, p+1, high, k)
	}
}

// 由大到小排序
func partitionSortDesc(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] >= pivot {
			//i为慢指针，记录小于基准值的元素位置索引
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	//确定基准值的位置并置换
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

//获取无序数组最小的前k个数
//可用优先队列解决，也可用快排思想
// 对数组[l, r]一次快排partition过程可得到，[l, p), p, [p+1, r)三个区间,[l,p)为小于等于p的值
// [p+1,r)为大于等于p的值。
// 然后再判断p，利用二分法
//     如果[l,p), p，也就是p+1个元素（因为下标从0开始），如果p+1 == k, 找到答案
//     2。 如果p+1 < k, 说明答案在[p+1, r)区间内，
//     3， 如果p+1 > k , 说明答案在[l, p)内
func findLastK(arr []int, k int) []int {
	res := []int{}
	if k == 0 || k > len(arr) {
		return res
	}
	low, high := 0, len(arr)-1
	for low < high {
		p := partitionSort(arr, low, high)
		if p+1 == k {
			return arr[:k]
		}
		if p+1 < k {
			low = p + 1
		} else {
			high = p - 1
		}
	}
	return res
}

//使用最小优先队列解决,数组容量必须大于2k且后边有足够多的小数可以替换队列中的大数
func findLastKWithQueue(arr []int, k int) []int {
	queue := []int{}
	if k == 0 || k > len(arr) {
		return queue
	}
	for _, v := range arr {
		if len(queue) < k {
			//该队列必须保证由大到小有序
			queue = append(queue, v)
			// if v > queue[0] {
			// 	queue[0], queue[i] = queue[i], queue[0]
			// }
		} else {
			if v < queue[0] {
				queue = queue[1:]
				queue = append(queue, v)
			}
		}
	}
	return queue
}

func TestFindLastK(t *testing.T) {
	arr := []int{5, 3, 6, 9, 2, 1, 4, 10, 11, 7, 9, 15}
	res := findLastK(arr, 5)
	// res := findLastKWithQueue(arr, 5)
	t.Log(res)
}
