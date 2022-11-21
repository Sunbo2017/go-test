package model

import (
	"fmt"
	"sync"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func (n *ListNode) Show() {
	fmt.Printf("val:%v -> next:%+v\n", n.Val, n.Next)
	for n.Next != nil {
		n = n.Next
		fmt.Printf("val:%v -> next:%+v\n", n.Val, n.Next)
	}
}

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
