package test

import (
	"fmt"
	"strings"
	"testing"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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
		result = max(result, right-left)
		// fmt.Println(s[left:right])
	}
	return result
}

func TestLongestSubstring(t *testing.T) {
	// s := "pwwkew"
	s := "abccbb"
	l := longestSubstring(s)
	fmt.Println(l)
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

	for i := 0; i < len1-len2+1; i++{
		if s1[i:i+len2] == s2{
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
