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

func TestValid(t *testing.T) {
	s1 := "()[]{}"
	s2 := "([]){}"
	s3 := "([)]{}"
	fmt.Println(isValid(s1))
	fmt.Println(isValidSample(s2))
	fmt.Println(isValid(s3))
}
