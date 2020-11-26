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
func longestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 扩展 ASCII 码的位图表示（BitSet），共有 256 位
	var bitSet [256]uint8
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && bitSet[s[right]] == 0 {
			// 右侧已经检查过的字符置为1
			bitSet[s[right]] = 1
			right++
		} else {
			//将已比较过不满足要求的位值归0，开始下一轮查找
			bitSet[s[left]] = 0
			left++
		}
		result = max(result, right-left)
		fmt.Println(s[left:right])
	}
	return result
}

func TestLongestSubstring(t *testing.T) {
	s := "pwwkew"
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

func TestStrStr(t *testing.T){
	s1 := "GodBlessYou"
	s2 := "less"
	fmt.Println(strings.Index(s1, s2))
	fmt.Println(strStr(s1, s2))
}
