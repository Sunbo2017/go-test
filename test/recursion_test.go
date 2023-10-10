package test

import (
	"fmt"
	"testing"
)

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

//棋盘路径问题
//一个机器人在m×n大小的地图的左下角，坐标(0,0)。
//机器人每次可以向上或向右移动。机器人要到达地图的右上角（终点）。
//可以有多少种不同的路径从起点走到终点？
//到达(x,y)处，可以由(x-1,y)处向右走一步，也可以由(x,y-1)处向上走一步
//极限场景：只有一行(y=0)或一列(x=0)情况,只有一种走法
func pathNum(x, y int) int {
	//如果路径坐标为交叉点位，判断条件为x == 0 || y == 0
	//如果路径坐标为格子点位，判断条件为x == 1 || y == 1
	if x == 0 || y == 0 {
		return 1
	}

	return pathNum(x, y-1) + pathNum(x-1, y)
}

func TestPath(t *testing.T) {
	t.Log(pathNum(2, 2))
}
