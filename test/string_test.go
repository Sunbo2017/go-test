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
	return result
}

//使用map替换bitmap
func longestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 
	bitSet := make(map[byte]int, 256)
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
    return s[left+1:right]
}
