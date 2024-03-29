package test

import (
	"fmt"
	"strings"
	"testing"
)

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

// LeetCode-cookbook-3： 查找最大不重复子串，返回其长度
// 位图法：位图就是bitmap的缩写，所谓bitmap，是用每一个bit位来存放某种状态，
// 位图适用于大规模数据，但数据状态又不是很多的情况。通常是用来判断某个数据存不存在的。
// 滑动窗口思想
func longestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 扩展 ASCII 码的位图表示（BitSet），共有 256 位
	var bitSet [256]uint8
	// 结果，窗口起点和终点
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && bitSet[s[right]] == 0 {
			// 右侧已经检查过的字符置为1
			bitSet[s[right]] = 1
			right++
		} else {
			//将已比较过的left位值归0，开始下一轮查找
			bitSet[s[left]] = 0
			fmt.Println(s[left:right])
			left++
		}
		result = Max(result, right-left)
		// fmt.Println(s[left:right])
	}
	fmt.Println(s[left:right])
	return result
}

//使用map替换bitmap
func longestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	//
	bitSet := make(map[byte]byte, 256)
	// 结果，窗口起点和终点
	result, left, right := 0, 0, 0
	//右指针遇到重复值左指针就会接着走并将原来位置归0，所以左右指针之间一定是不重复子串，右指针走到头时，左指针也没必要继续走
	for right < len(s) {
		if bitSet[s[right]] == 0 {
			// 右侧已经检查过的字符置为1
			bitSet[s[right]] = 1
			right++
		} else {
			//将已比较过的left位值归0，开始下一轮查找，直到左指针走过了相等的元素后，右指针接着往下走
			bitSet[s[left]] = 0
			fmt.Println(s[left:right])
			left++
		}
		result = Max(result, right-left)
		// fmt.Println(s[left:right])
	}
	fmt.Println(s[left:right])
	return result
}

func TestLongestSubstring(t *testing.T) {
	// s := "pwwkew"
	// s := "abrfccbde"
	//s := "aaaccbde"
	s := "abcdeeabde"
	l1 := longestSubstring(s)
	fmt.Println(l1)
	fmt.Println("---------")
	l2 := longestSubstring1(s)
	fmt.Println(l2)
}

// 查找 substring,如果在母串中找到了子串，返回子串在母串中出现的下标，
// 如果没有找到，返回 -1，如果子串是空串，则返回 0 。
func strStr(s1, s2 string) int {
	// return strings.Index(s1, s2)
	if len(s2) == 0 {
		return 0
	}
	if len(s1) < len(s2) {
		return -1
	}
	if s1 == s2 {
		return 0
	}
	// 遍历s1
	for i := 0; ; i++ {
		// 遍历s2
		for j := 0; ; j++ {
			if j == len(s2) {
				return i
			}
			if i+j == len(s1) {
				return -1
			}
			if s2[j] != s1[i+j] {
				break
			}
		}
	}
}

func strStr1(s1, s2 string) int {
	len1 := len(s1)
	len2 := len(s2)

	if len2 == 0 {
		return 0
	}
	if len1 < len2 {
		return -1
	}
	if s1 == s2 {
		return 0
	}

	for i := 0; i < len1-len2+1; i++ {
		if s1[i:i+len2] == s2 {
			return i
		}
	}
	return -1
}

func TestStrStr(t *testing.T) {
	s1 := "GodBlessYou"
	s2 := "You"
	fmt.Println(strings.Index(s1, s2))
	fmt.Println(strStr(s1, s2))
	fmt.Println(strStr1(s1, s2))
}

func TestCompare(t *testing.T) {
	s1 := "1.6"
	s2 := "1.7"
	//直接比较时，从左至右逐个字符根据ASCII码值大小比较
	fmt.Println(s1 > s2) //flase
	// 如果待比较的字符串位数不等，则可能会出错
	fmt.Println("3" > "15") //true
}

//查找最长回文子串：从中间字符开始依次比较前后字符，若相等即为回文子串，不等则继续寻找
//中间位置可能是某个字符，如aba；也可能是空白，如aa
func longestPalindrome(s string) string {
	var sub, longest string
	var mid int

	for mid < len(s) {
		// 对称点为空白，长度为偶数
		sub = findLongestPalindromeByMid(s, mid-1, mid)
		if len(sub) > len(longest) {
			longest = sub
		}

		// 对称点为某个字符，长度为奇数
		sub = findLongestPalindromeByMid(s, mid-1, mid+1)
		if len(sub) > len(longest) {
			longest = sub
		}
		mid++
	}
	return longest
}

func findLongestPalindromeByMid(s string, left, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	//注意此处,left索引位置不相等后退出循环，所以相等的位置是left+1和right-1，即[left+1,right),左闭右开
	return s[left+1 : right]
}

func TestLongest(t *testing.T) {
	str := "abcdefedca"
	subS := longestPalindrome(str)
	t.Log(subS)
}

//金山云面试题：字符串加法：需使用纯字符串，不可转为int类型
func AddStrNumber(a, b string) string {
	l1, l2, max := len(a), len(b), 0
	if l1 >= l2 {
		max = l1
	} else {
		max = l2
	}
	//存储逆序的结果
	var res []string
	num1, num2 := make([]byte, max), make([]byte, max)

	c1, c2 := 0, 0
	for i := l1 - 1; i >= 0; i-- {
		//逆序存储字符串a每个字符，若a长度小于max，默认0补齐
		// num1 = append(num1, a[i]-'0')
		num1[c1] = a[i] - '0'
		c1++
	}
	for i := l2 - 1; i >= 0; i-- {
		//逆序存储字符串b每个字符，若b长度小于max，默认0补齐
		// num2 = append(num2, b[i]-'0')
		num2[c2] = b[i] - '0'
		c2++
	}
	//保存进位值
	var up byte = 0
	for i := 0; i < max; i++ {
		//注意先计算和，再更新进位值
		sum := (up + num1[i] + num2[i]) % 10
		up = (up + num1[i] + num2[i]) / 10
		res = append(res, string(sum+'0'))
	}
	result := ""
	c := len(res) - 1
	//逆序拼接数组元素获得正序的和
	for i := len(res) - 1; i >= 0; i-- {
		//去除开头的0
		if c == i && res[i] == "0" {
			c--
			continue
		}
		result += res[i]
	}
	return result
}

func AddBigStrNumber(a, b string) string {
	result := ""
	if len(a) == 0 && len(b) == 0 {
		return "0"
	}
	var index1 = len(a) - 1
	var index2 = len(b) - 1
	var up int

	for index1 >= 0 || index2 >= 0 {
		c1, c2 := 0, 0
		if index1 >= 0 {
			c1 = int(a[index1] - '0')
		}
		if index2 >= 0 {
			c2 = int(b[index2] - '0')
		}

		sum := c1 + c2 + up
		up = sum / 10
		c3 := (sum % 10) + '0'

		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}
	//处理最后的进位值
	if up == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	start := 0
	//移除开头的0
	for i, v := range result {
		if v-'0' != 0 {
			start = i
			break
		}
	}
	return result[start:]
}

func TestAddStrNumber(t *testing.T) {
	a, b := '2'-'0', '3'-'0'
	t.Log('0')
	t.Log(a)
	t.Log(b)
	t.Log(a + b)

	str1, str2 := "001234587669003", "00034579" // res=1234587703582
	res := AddStrNumber(str1, str2)
	t.Log(res)

	res1 := AddBigStrNumber(str1, str2)
	t.Log(res1)

	str1, str2 = "1234587669003", "9000000034579"
	res2 := AddBigStrNumber(str1, str2)
	t.Log(res2)
}

//字符串乘法
func MultiStrNumber(a, b string) string {
	l1, l2 := len(a), len(b)

	var res [][]byte
	//进位值
	up := byte(0)
	for i := l1 - 1; i >= 0; i-- {
		n1 := a[i] - '0'
		//逆序存储每一位
		r := []byte{}
		for j := l2 - 1; j >= 0; j-- {
			n2 := b[j] - '0'
			sum := (n1*n2 + up)
			cur := sum % 10
			up = sum / 10
			r = append(r, cur)
		}
		if up > 0 {
			r = append(r, up)
		}
		reverse := reverseArray(r)
		c := l1 - 1 - i
		//从第二轮乘积开始在末位补0
		if c >= 0 {
			add := make([]byte, c)
			reverse = append(reverse, add...)
		}
		res = append(res, reverse)
	}
	temp := byteArray2str(res[0])
	for i := 1; i < len(res); i++ {
		cur := byteArray2str(res[i])
		temp = AddStrNumber(cur, temp)
	}
	return temp
}

func reverseArray(arr []byte) []byte {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func byteArray2str(arr []byte) string {
	res := ""
	for _, v := range arr {
		res += string(v + '0')
	}
	return res
}

func TestMultiStrNumber(t *testing.T) {
	a := "123123"
	b := "20000111"
	res := MultiStrNumber(a, b)
	t.Log(res)
}

//判定子序列,如下：
//s="abc", t="ahbgdc", return true
//s = "axc", t = "ahbgdc", return false.
//利⽤双指针  i, j  分别指向  s, t  ，⼀边前进⼀边匹配⼦序列
func judgeSubStr(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

// 定义重复字符串是由两个相同的字符串首尾拼接而成。例如："abcabc" 是一个长度为 6 的重复字符串，
// 因为它由两个 "abc" 串拼接而成；"abcba" 不是重复字符串，因为它不能由两个相同的字符串拼接而成。
// 给定一个字符串，请返回其最长重复子串的长度。
// 若不存在任何重复字符子串，则返回 0。
func maxRepeatSubstr(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); j++ {
			sub := s[i:j]
			if len(sub)%2 == 0 {
				sub1 := sub[:len(sub)/2]
				sub2 := sub[len(sub)/2:]
				if sub1 == sub2 {
					max = Max(max, len(sub))
				}
			}
		}
	}
	return max
}

// func maxRepeatSubStr1(s string) int {
// 	max := 0
// 	mid := len(s)/2
// 	//从中间向左比较重复子串
// 	for i:=mid;i>0;i-- {
// 		if s[:mid+1] == s[mid+1:(mid+1)*2] {
// 			max = (mid+1) * 2
// 			break
// 		}
// 	}

// 	//从中间向右
// 	for i:=mid+1;
// 	return max
// }

func TestMaxRepeat(t *testing.T) {
	// s1 := "abcabcdd"
	// t.Log(maxRepeatSubstr(s1))
	// s2 := "abcab"
	// t.Log(maxRepeatSubstr(s2))
	s3 := "ababcdefcdef"
	t.Log(maxRepeatSubstr(s3))
	// t.Log(maxRepeatSubStr1(s3))
}

//查找两个字符串最长公共子串
func LCS(str1 string, str2 string) string {
	// 二维数组存储所有的str1和str2相等字符的索引位置
	var dp [][]int
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				//有连续相等的子串，则不用添加后续索引位置
				if len(dp) > 0 {
					temp := dp[len(dp)-1]
					if i-temp[0] == j-temp[1] {
						continue
					}
				}
				in := []int{i, j}
				dp = append(dp, in)
			}
		}
	}
	res := ""
	for _, v := range dp {
		i, j := v[0], v[1]
		for i < len(str1) && j < len(str2) {
			if str1[i] == str2[j] {
				if i-v[0] >= len(res) {
					res = str1[v[0] : i+1]
				}
				i++
				j++
			} else {
				break
			}
		}
	}
	return res
}

func TestLCS(t *testing.T) {
	a := "abcgfhjkdsd"
	b := "jfghthjkd"

	res := LCS(a, b)
	t.Log(res)
}

//最小覆盖字串
//输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
func minWindow(s string, t string) string {
	wind := make(map[byte]int)
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	left, right, match, start, end, min := 0, 0, 0, 0, 0, len(s)
	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			wind[c]++
			if wind[c] == need[c] {
				match++
			}
		}
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			left++
			if need[c] != 0 {
				//可能存在s-->aaaa  t-->a,
				//这里只有当s的最后一个a也被移出窗口的时候
				//匹配数才少了1
				if wind[c] == need[c] {
					match--
				}
				wind[c]--
			}
		}
	}
	if min == len(s) {
		return ""
	}
	return s[start:end]
}

//给定一个字符串，按照a->z,b->y...m->n,n->m规则置换
func replaceStr(s string) string {
	var resBytes []byte
	for _, v := range s {
		if v <= 'm' {
			c := 'm' - v
			n := 'n' + c
			resBytes = append(resBytes, byte(n))
		} else if v >= 'n' {
			c := v - 'n'
			m := 'm' - c
			resBytes = append(resBytes, byte(m))
		}
	}

	return string(resBytes)
}

func TestReplaceStr(t *testing.T) {
	s := "abcmnxyz"
	r := replaceStr(s)
	t.Log(r)
}

//采用二维DP的方法。定义一个dp表,其中dp[i][j]表示str1[0..i-1]和str2[0..j-1]的最长公共子序列长度。递推方程如下:
//如果str1[i-1]==str2[j-1],那么dp[i][j] = dp[i-1][j-1] + 1
//如果str1[i-1]!=str2[j-1],那么dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//最终dp[m][n]就是str1和str2的最长公共子序列长度。
//时间复杂度O(MN),空间复杂度O(MN)。
func longestCommonSubsequence(str1 string, str2 string) int {
	m := len(str1)
	n := len(str2)

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	if dp[m][n] == 0 {
		return 0
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
