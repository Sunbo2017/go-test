package test

import (
	"testing"
)

// 某商店规定：三个空汽水瓶可以换一瓶汽水，允许向老板借空汽水瓶（但是必须要归还）。
// 小张手上有n个空汽水瓶，她想知道自己最多可以喝到多少瓶汽水。
func maxBottles(n int) int {
	if n == 1 {
		return 0
	} else if n==2 {
		return 1
	} else {
		a := n/3
		b := n%3
		return a + maxBottles(a+b)
	}
}

func TestBottles(t *testing.T) {
	n := 81
	s := maxBottles(n)
	t.Log(s)
}