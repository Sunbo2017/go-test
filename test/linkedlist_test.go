package test

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Show() {
    fmt.Printf("%v :: %v\n", n.Val, n.Next)
    for n.Next != nil {
        n = n.Next
		fmt.Printf("%v :: %v\n", n.Val, n.Next)
    }
}

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{0, nil}
	//记录链表尾节点
	tail := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return head.Next
}

func TestMergeTwoLists(t *testing.T) {
	l1, l2 := &ListNode{0, nil}, &ListNode{0, nil}
	cur1, cur2 := l1, l2
	for i := 0; i < 5; i++ {
		cur1.Next = &ListNode{2*i, nil}
		cur1 = cur1.Next
	}
	fmt.Println("l1:")
	l1.Next.Show()
	for i := 0; i < 4; i++ {
		cur2.Next = &ListNode{3*i, nil}
		cur2 = cur2.Next
	}
	fmt.Println("l2:")
	l2.Next.Show()

	fmt.Println("merge:")
	mergeList := mergeTwoLists(l1.Next, l2.Next)
	mergeList.Show()
}
