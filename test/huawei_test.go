package test

import (
	// "strings"
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


//种一排树共n棵，其中有m棵未能成活，给定一个最多可补种的数量t，
//和一个数组a，表示未成活的序号，t<=len(a)=m<n,求如何补种可获得最长的序列，返回最长序列长度
// func maxTreeList(a []int, n, t int) int {
// 	if t == len(a) {
// 		return n
// 	}
// 	s := make([]string, n)
// 	for i:=0;i<n;i++ {
// 		s[i]="1"
// 	}
// 	for _, v := range a {
// 		//表示此棵树未成活
// 		s[v] ="0" 
// 	}
// 	line := strings.Join(s,"")

// 	treeMap := make(map[int][][]int, 0)
// 	for i:=1;i<=t;i++ {
// 		arr1 := [][]int{}
// 		for j,v := range a {
// 			arr2 := make([]int, len(a))
			
// 		}
// 	}

// 	for i:=1;i<=t;i++ {
// 		c := 0
// 		for _, v := range a {
// 			c++
// 			s[v]="1"

// 		}
// 	}


// 	return 0
// }