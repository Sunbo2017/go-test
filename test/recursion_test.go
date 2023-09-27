package test

import "fmt"

//递归求解斐波那契数列
func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

//求阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

//倒序输出一个正整数
func printDigit(n int) {
	fmt.Println(n % 10)
	if n > 10 {
		printDigit(n / 10)
	}
}

//字节一面
//一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法
//递归
func jumpFloor1(N int) int {
	if N <= 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return N
	}
	return jumpFloor1(N-1) + jumpFloor1(N-2)
}

//动态规划
func jumpFloor2(N int) int {
	if N <= 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return N
	}
	a, b := 1, 2
	for i := 3; i <= N; i++ {
		a, b = b, a+b
	}
	return b
}

// 使用dpTable
func jumpFloor3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	//存储每一步的结果
	dpTable := make([]int, n+1)

	// base case
	dpTable[1] = 1
	dpTable[2] = 2

	// 状态转移
	for i := 3; i <= n; i++ {
		dpTable[i] = dpTable[i-1] + dpTable[i-2]
	}
	return dpTable[n]
}
