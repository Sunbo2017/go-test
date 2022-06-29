package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func testMaxBottles() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "0" {
			return
		}
		n, _ := strconv.Atoi(s.Text())
		fmt.Println(maxBottles(n))
	}
}

func testAdd () {
	a, b := 0, 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		}else {
			fmt.Println(a+b)
		}
	}
}

func testAdd1() {
	t, a, b := 0, 0, 0
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		}
		fmt.Println(a+b)
	}
}

func testAdd2() {
	a, b := 0, 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		}
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a+b)
	}
}

func testAdd3() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		arr := strings.Split(s.Text(), " ")
		if arr[0] == "0" {
			return
		}
		sum := 0
		for i:=1;i<len(arr);i++ {
			t, _ := strconv.Atoi(arr[i])
			sum += t
		}
		fmt.Println(sum)
	}
}

func testAdd4() {
	var n, c int
    fmt.Scan(&n)
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        t := s.Text()
		fmt.Printf("t===%v\n", t)
        strs := strings.Split(t," ")
        if len(strs) == 1 {
			fmt.Println("continue...")
            continue
        }
		l, _ := strconv.Atoi(strs[0])
		if l != len(strs) -1 {
			fmt.Println("error")
			return
		}
        var sum int
        for i:=1;i<len(strs);i++ {
            v,_ := strconv.Atoi(strs[i])
            sum += v
        }
        fmt.Println(sum)
        c++
		if c == n {
			return
		}
    }
}

func testSort() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	fmt.Println(n)
	//取得第二行数据
	s.Scan()
	arr := strings.Split(s.Text(), " ")
	fmt.Println(arr)
	if len(arr) != n {
		fmt.Println("error")
		return
	}
	sort.Strings(arr)
	fmt.Println(strings.Join(arr, " "))
}

func testSort1() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		arr := strings.Split(s.Text(), " ")
		sort.Strings(arr)
		fmt.Println(strings.Join(arr, " "))
	}
}


func getOffset() []int {
	offset := make([]int, 50)
	offset[0],offset[1],offset[2]=1,2,4
	for i:=3;i<50;i++ {
		offset[i] = offset[i-1]+offset[i-2]+offset[i-3]
	}
	return offset
} 

func strEncode() {
	line := 0
	n, _ := fmt.Scan(&line)
	if n == 0 {
		return
	}
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	offset := getOffset()
	for s.Scan() {
		res := []byte{}
		arr := s.Text()
		fmt.Println(arr)
		for i, a := range arr {
			offsetI := offset[i]
			fmt.Printf("off===%v\n",offsetI)
			var v int64 = int64(a) + int64(offsetI)
			if v > int64('z') {
				v = (v-int64('z'))%26 + int64('a') - 1
			}
			res = append(res, byte(v))
		}
		fmt.Println(string(res))
	}
}






func main() {
	// testAdd()
	// testAdd1()
	// testAdd2()
	// testAdd3()
	// testAdd4()
	// testSort()
	// testSort1()
	strEncode()
}