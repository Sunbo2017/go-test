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

// LeetCode-cookbook-21：合并两个有序链表
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
		cur1.Next = &ListNode{2 * i, nil}
		cur1 = cur1.Next
	}
	fmt.Println("l1:")
	l1.Next.Show()
	for i := 0; i < 4; i++ {
		cur2.Next = &ListNode{3 * i, nil}
		cur2 = cur2.Next
	}
	fmt.Println("l2:")
	l2.Next.Show()

	fmt.Println("merge:")
	// mergeList := mergeTwoLists(l1.Next, l2.Next)
	mergeList := addTwoNumbers(l1.Next, l2.Next)
	mergeList.Show()
}

//LeetCode-cookbook-2：两个逆序链表相加
//Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
//Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	//虚拟头结点，可理解为头指针
	head := &ListNode{Val: 0, Next: nil}
	//相当于尾节点，应用尾插法
	current := head
	//进位，初始值为0
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}
		//取余求当前位的值
		current.Next = &ListNode{Val: (x + y + carry) % 10, Next: nil}
		current = current.Next
		//求进位值（加法计算最大只能为1）
		carry = (x + y + carry) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	//判断最后是否需要进位
	if carry > 0 {
		current.Next = &ListNode{Val: carry % 10, Next: nil}
	}
	return head.Next
}

// 单链表快排
func quickSortList(start, end *ListNode) {
	if start == end || start.Next == end {
		return
	}
	mid := partion(start, end)
	quickSortList(start, mid)
	quickSortList(mid.Next, end)
}

func partion(start, end *ListNode) *ListNode {
	if start == end || start.Next == end {
		return start
	}

	pivot := start.Val //选择基准
	p := start
	q := start
	for q != end {
		if q.Val < pivot {
			p = p.Next
			p.Val, q.Val = q.Val, p.Val
		}
		q = q.Next //否则一直往下走
	}
	
	p.Val, start.Val = start.Val, p.Val //定位
	return p
}
