package test

import (
	"fmt"
	"sort"
	"testing"
)

// map append问题
func TransformResultFormat(columns []string, values [][]string) {

	var result []map[string]interface{}
	var TransformResult = make(map[string]interface{})

	fmt.Println("yyyy", len(values))
	for i := 0; i < len(values); i++ {
		for j, v := range columns {
			TransformResult[v] = values[i][j]
		}
		fmt.Println("xxx", TransformResult)
		result = append(result, TransformResult)
	}
	//result = append(result, TransformResult)

	fmt.Println(result)
}

func TestMapTransfor(t *testing.T) {
	c := []string{"name", "grade", "uuid"}
	v := [][]string{
		{"xiaoming", "5", "59525F6C427F339F88B5C81FE9DC3671"},
		{"xiaoli", "51", "D08594DF1F983809A890CF024E64B06B"},
	}
	TransformResultFormat(c, v)
}

// 字节面试题：台阶积水问题
// 忽略第0个台阶，从第1个台阶开始算，后边台阶如果矮于第一个台阶，差值即为积水量，
// 直到高于第一个台阶的新台阶出现，第一次循环结束，
// 然后从这个新台阶开始第二次循环
func TestStepsWater(t *testing.T) {
	// steps := []int{1, 0, 2, 1, 2, 3, 0, 1, 2, 4}
	steps := []int{0, 0, 2, 1, 2, 3, 0, 1, 2, 4}
	water := 0
	i, j := 0, 1
	for i < len(steps) {
		tempWater := water
		for i+j < len(steps) && steps[i+j] <= steps[i] {
			water += steps[i] - steps[i+j]
			j++
		}
		// 如果一直到最后一个台阶也不高于起始台阶，则无法积水
		if i+j == len(steps)-1 && steps[i+j] <= steps[i] {
			water = tempWater
		}
		// 直接跳到新台阶进行下一次循环
		i += j
		j = 1
	}
	fmt.Println(water)
}

func TestSort(t *testing.T) {
	ints := []int{0, 5, 2, 1, 3, 4, 6, 9, 8, 7}
	sort.Ints(ints)
	fmt.Println(ints)
}

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res = []string{}
)

// LeetCode-cookbook-17：根据手机数字按钮返回对应数字可生成的所有字母组合
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}
func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}

func TestCombinations(t *testing.T) {
	letterCombinations("456")
	fmt.Println(res)
	byten := byte('4')
	r := byten - '1'
	fmt.Println(r)
}
