package test

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

//二分法求平方根
func mySqrt(n int) int {
	l, r := 1, n
	for {
		middle := (l + r) / 2
		if middle <= n/middle && (middle+1) > n/(middle+1) {
			return middle
		} else if middle < n/middle {
			l = middle + 1
		} else {
			r = middle - 1
		}
	}
	return 0
}

func TestSqrt(t *testing.T) {
	n := 10
	s := math.Sqrt(float64(n))
	fmt.Println(s)

	s1 := mySqrt(n)
	fmt.Println(s1)
}

func deferReturn() int {
	i := 0
	defer func() {
		i += 1
		fmt.Println("defer1", i)
	}()
	defer func() {
		i += 1
		fmt.Println("defer2", i)
	}()
	//panic(i)
	return i
}

func TestDefer(t *testing.T) {
	t.Log(deferReturn())
}

func copyArray() {
	var test, another []uint8
	fmt.Println("len1:", len(test), cap(test))
	//0  0

	test = make([]uint8, 5, 10)
	fmt.Println("len2:", len(test), cap(test))
	//5  10

	another = append(test, 1, 2, 3)
	fmt.Println("len3:", len(another), cap(another))
	//8  10

	fmt.Println("len4:", len(test), cap(test))
	//5  10

	copy(test, []uint8{6, 6, 6, 6})
	fmt.Println("content1:", another, test)

	another = append(another, 1, 2, 3)
	fmt.Println("len5:", len(another), cap(another))

	copy(test, []uint8{5, 5})
	fmt.Println("content2:", another[0], test[0])

	ch := make(chan int, 5)
	fmt.Println("len6:", len(ch), cap(ch))

	ch <- 1
	ch <- 2
	fmt.Println("len7:", len(ch), cap(ch))

	type s1 struct {
		p1 uint8
		p2 uint32
		p3 uint64
	}
	var v1 s1
	fmt.Println("len8:", unsafe.Sizeof(v1.p1), unsafe.Sizeof(v1.p2), unsafe.Sizeof(v1.p3), unsafe.Sizeof(v1))
}

func TestCopy(t *testing.T) {
	copyArray()
}

//使用iota定义下面的枚举值
//a1=1 a2=2 a3=7 a4=16
func TestIota(t *testing.T) {
	//const每新增一行常量声明，iota计数一次，可以当做const语句中的索引，常用于定义枚举数据。
	const (
		a = iota
		a1
		a2
		a3 = iota + 4
		a4 = iota + 12
	)

	t.Log(a1)
	t.Log(a2)
	t.Log(a3)
	t.Log(a4)
}
