package test

import (
	"fmt"
	"math"
	"testing"
)

//二分发求平方根
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
