package test

//输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//双重循环暴力求解
func findMaxSubArray(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			max = Max(max, sum)
		}
	}
	return max
}

//动态规划，使用dp table,dp[i]存储以元素arr[i]为结尾的子数组之和的最大值
//状态转移方程： dp[i] = Math.max(dp[i-1]+arr[i], arr[i]);
func findMaxSubArrByDp(arr []int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		dp[i] = Max(dp[i-1]+arr[i], arr[i])
		max = Max(dp[i], max)
	}
	return max
}

//动态规划，不使用dp table,直接用sum存储以元素arr[i]为结尾的子数组之和的最大值
//状态转移方程： dp[i] = Math.max(dp[i-1]+arr[i], arr[i]);
func findMaxSubArr(arr []int) int {
	sum := 0
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		sum = Max(sum+arr[i], arr[i])
		max = Max(sum, max)
	}
	return max
}
