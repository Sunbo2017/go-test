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

// 字节面试题：台阶积水问题，类似滑动窗口思想
// 忽略第0个台阶，从第1个台阶开始算，后边台阶如果矮于第一个台阶，差值即为积水量，
// 直到高于第一个台阶的新台阶出现，第一次循环结束，
// 然后从这个新台阶开始第二次循环
// 该思想存在局限性，不能当作正确答案
func TestStepsWater(t *testing.T) {
	// steps := []int{1, 0, 2, 1, 2, 3, 0, 1, 2, 4}
	steps := []int{0, 0, 2, 1, 2, 3, 0, 1, 3, 2}
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

type step struct {
	val int  //当前台阶高度
	leftMax int  //左侧最高台阶高度
	rightMax int  //右侧最高台阶高度
}

func makeSteps(steps []int) []step {
	list := make([]step, len(steps))
	for i, v := range steps {
		s := step{val: v}
		list[i] = s
	}
	for i:=1; i<len(steps)-1; i++ {
		list[i].leftMax = Max(list[i-1].leftMax, list[i-1].val)
	}
	for i:=len(steps)-2; i>0; i-- {
		list[i].rightMax = Max(list[i+1].rightMax, list[i+1].val)
	}
	return list
}

func TestStepWater1(t *testing.T) {
	steps := []int{0, 0, 2, 1, 2, 3, 0, 1, 3, 2}
	water := 0	
	stepList := makeSteps(steps)
	for i:=1; i<len(steps)-1; i++ {
		increment := Min(stepList[i].leftMax - stepList[i].val, stepList[i].rightMax - stepList[i].val)
		if increment > 0 {
			water += increment
		}
	}
	t.Logf("water:%v", water)
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

// var res [][]string

// 已知每一个字母可以用 1~26 表示，给定一个数字组成的字符串，问它可以表示多少种字母组合。
// 123456215
func numString(str string, lStr []string) int {
	for i:=0; i<len(str); {

	}
	
	return 0
}

// 输入一个数组和一个目标值 T，判断数组中是否存在两个数的和为 T
// 最简单方法可以直接双层循环判断和，O(n^2)
// 可以使用map的k，v分别记录元素值和差值，O（n）
func judege2Sum(arr []int, t int) (v1,v2 int){
	resMap := map[int]int{}
	for _, v := range arr{
		if _, ok := resMap[v]; !ok{
			resMap[t - v] = v
		} else{
			return v, resMap[v]
		}
	}
	return 0, 0
}

// 升级：输入一个数组和一个目标值 T，判断数组中是否存在某些数的和为 T。



func Test2Sum(t *testing.T){
	array := []int{1,2,3,4,6,5,8,9}
	// 6, 4
	v1, v2 := judege2Sum(array, 10)
	fmt.Println(v1)
	fmt.Println(v2)
}

// 一个细胞的寿命是5min 他会在2min和4min 分别分裂出一个新细胞，请问n min后 ，有多少细胞 
var sumn = 1
func sum(n int) int{
    for i:=1;i<=n;i++{
        if i%2 == 0 || i%4==0 {
            // sumn += sum(n-i)
			sumn *= 2
        }
        if i%5==0 {
			div := i/5
            sumn = sumn - 2*div
        }
    }
    return sumn
}

func TestSum(t *testing.T){
	n := 15
	fmt.Println(sum(n))
}