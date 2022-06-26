package test

import (
	"fmt"
	"sync"
	"testing"
)

type (
	node struct {
		val  interface{}
		prev *node
	}

	Stack struct {
		top    *node
		length int
		lock   *sync.Mutex
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0, &sync.Mutex{}}
}

func (s *Stack) Push(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top = &node{val, s.top}
	s.length++
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.val
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.val
}

func (s *Stack) Len() int {
	return s.length
}

func Benchmark_Push(b *testing.B) {
	stack := NewStack()
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Push("test")
	}
}

func Benchmark_Pop(b *testing.B) {
	stack := NewStack()
	b.StopTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Push("test")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Pop()
	}
}

// LeetCode-cookbook-20：校验字符串是否是合规括号对
//Input: "{[]}" Output: true
//Input: "()[]{}" Output: true
//Input: "([)]" Output: false
// 使用结构体实现栈结构
func isValid(s string) bool {
	// 括号对字典
	bracketsMap := map[uint8]uint8{'{': '}', '[': ']', '(': ')'}
	// 传入字符串为空则返回true
	if s == "" {
		return true
	}
	// 初始化栈
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		// 如果栈中有数据，则进行比对，如果栈顶元素和当前元素匹配，则弹出，其他情况向栈中压入元素
		if stack.Len() > 0 {
			if c, ok := bracketsMap[stack.Peek().(uint8)]; ok && c == s[i] {
				stack.Pop()
				continue
			}
		}
		// 后进行入栈操作，确保栈顶元素和当前字符能进行括号对比较
		stack.Push(s[i])
	}
	// 到最后如果栈不为空，则说明未完全匹配掉（完全匹配的话，栈只能为空）
	return stack.Len() == 0
}

// LeetCode-cookbook-20：校验字符串是否是合规括号对
// 直接使用切片实现栈结构
func isValidSample(s string) bool {
	// 括号对字典
	bracketsMap := map[rune]rune{'{': '}', '[': ']', '(': ')'}
	// 空字符串直接返回 true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if c, ok := bracketsMap[stack[len(stack)-1]]; ok && c == v && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

// LeetCode 921：使括号有效的最少添加
// 给你输⼊⼀个字符串 s，你可以在其中的任意位置插⼊左括号 ( 或者右括号 )，
// 请问你最少需要⼏次插⼊才能使得 s 变成⼀个有效的括号串？
// ⽐如说输⼊ s = "())("，算法应该返回 2，因为我们⾄少需要插⼊两次把 s 变成 "(())()"，这样每个左
// 括号都有⼀个右括号匹配，s 是⼀个有效的括号串
func minAdd2MakeValid(s string) int {
	// 分别记录左括号和右括号的需求数
	left, right := 0, 0
	for _, v := range s {
		if v == '(' {
			// 对右括号的需求 + 1
			right++
		}
		if v == ')' {
			// 对右括号的需求 - 1
			right--
			if right == -1 {
				right = 0
				// 需插⼊⼀个左括号
				left++
			}
		}
	}
	return left+right
}

func TestMinAdd(t *testing.T) {
	s := "())("
	count := minAdd2MakeValid(s)
	t.Log(count)
}

func TestValid(t *testing.T) {
	s1 := "()[]{}"
	s2 := "([]){}"
	s3 := "([)]{}"
	fmt.Println(isValid(s1))
	fmt.Println(isValidSample(s2))
	fmt.Println(isValid(s3))
}


// 单调栈:单调栈实际上就是栈，只是利⽤了⼀些巧妙的逻辑，使得每次新元素⼊栈后，栈内的元素都保持有序（单调递增或单调递减）。
// 给你⼀个数组 nums，请你返回⼀个等⻓的结果数组，结果数组中对应索引存储着下⼀个更⼤元素，如果没有更⼤的元素，就存 -1。
// ⽐如说，输⼊⼀个数组 nums = [2,1,2,4,3]，你返回数组 [4,2,4,-1,-1]。
// 解释：第⼀个 2 后⾯⽐ 2 ⼤的数是 4; 1 后⾯⽐ 1 ⼤的数是 2；第⼆个 2 后⾯⽐ 2 ⼤的数是 4; 
// 4 后⾯没有⽐ 4⼤的数，填 -1；3 后⾯没有⽐ 3 ⼤的数，填 -1。
func findNextGreater(nums []int) []int {
	res := make([]int, len(nums))
	s := NewStack()
	// 倒序入栈
	for i := len(nums)-1; i>=0; i-- {
		for s.Len() > 0 && s.Peek().(int) <= nums[i] {
			// 小于当前元素的值都出栈
			s.Pop()
		}
		if s.Len() > 0 {
			res[i] = s.Peek().(int)
		} else {
			res[i] = -1
		}
		s.Push(nums[i])
	}
	return res
}

func TestNextGreater(t *testing.T) {
	nums := []int{2,3,5,2,4,7,6}
	res := findNextGreater(nums)
	t.Log(res)
}




//直接使用数组实现最简单的栈结构
type Node struct {
	val interface{}
}

type NodeStack struct {
	Items []Node
}
 
 
func (n *NodeStack) New() *NodeStack {
	//
	n.Items = []Node{}
	return n
}
 
 
func (n *NodeStack) push(q Node) {
	n.Items = append(n.Items, q)
}
 
 
func (n *NodeStack) pop() *Node {
 
 
	item := n.Items[len(n.Items)-1]
 
 
	n.Items = n.Items[0: len(n.Items)-1]
	return &item
}
 
 
func (n *NodeStack) IsEmpty() bool {
	return len(n.Items) == 0
}
 
 
 func (n *NodeStack) Size() int {
	return len(n.Items)
 }