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
	done := make(chan int,5)
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
