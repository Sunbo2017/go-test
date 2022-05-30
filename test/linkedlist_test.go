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
	fmt.Printf("val:%v -> next:%+v\n", n.Val, n.Next)
	for n.Next != nil {
		n = n.Next
		fmt.Printf("val:%v -> next:%+v\n", n.Val, n.Next)
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


// 取出单链表倒数第k个元素：使用快慢指针
// 最简单粗暴方法遍历一遍链表，数据存到数组中，遍历结束后直接取数组倒数第k个值
// 使用两个距离相差k的指针，第一个node开始走时，k--，k=0时，第二个node开始走；
// 当第一个node走到结尾时，第二个node刚好位于倒数第k的位置，快指针和慢指针一直保持k个距离
func findBackwardsK(node *ListNode, k int) *ListNode {
	target := node
	current := node
	for current != nil {
		current = current.Next
		if k > 0 {
			k--
		} else {
			target = target.Next
		}
	}
	return target
}

// 移除单链表倒数第k个元素
// 注意边界问题，要移除倒数第k个节点，实际要找到倒数第k+1个节点
func rmBackwardsK(node *ListNode, k int) *ListNode {
	//k+1保证取到倒数第k个节点的前一个节点
	k += 1
	target := node
	current := node
	for current.Next != nil {
		current = current.Next
		k--
		if k <= 0 {
			target = target.Next
		}
		// if k > 0 {
		// 	k--
		// } else {
		// 	target = target.Next
		// }
	}
	//直接跳过倒数第k个节点指向下一个节点
	target.Next = target.Next.Next
	return node
}

func TestRmBackKList(t *testing.T) {
	head := &ListNode{0,nil}
	node := head
	for i := 1; i <= 10; i++ {
		node.Next = &ListNode{i, nil}
		node = node.Next
	}
	head.Next.Show()
	fmt.Println("---------------")
	res := rmBackwardsK(head.Next, 4)
	res.Show()
}


// 判断单链表是否有环：应用快慢指针
func judgeListCycle(node *ListNode) bool {
	fast, slow := node, node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 已知链表有环，找到环的起始位置
// 相遇时慢指针走k步，快指针走2k步，快指针比慢指针多走了k步，也就是环的长度为k
// 设相遇点距环的起点的距离为 m，那么环的起点距头结点 head 的距离为 k - m，
// 也就是说如果从 head 前进 k - m 步就能到达环起点。
// 巧的是，如果从相遇点继续前进 k - m 步，也恰好到达环起点。
func getCycleStart(node *ListNode) *ListNode {
	fast, slow := node, node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	slow = node
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}



// 翻转链表
// 从链表的第二个结点开始，把遍历到的结点插入到头结点的后面，直到遍历结束。
// 假如原链表为head->1->2->3->4->5->6->7，在遍历到2的时候，将2插入到头结点的后面，链表变为head->2->1->3->4->5->6->7，
// 同理head->3->2->1->4->5->6->7等等。
// 这种算法必须保证有虚拟头节点
func reverseLinkedlist(node *ListNode) {
	//node为头节点
	if node == nil || node.Next == nil {
		return
	}
	//暂存第二个节点
	cur := node.Next.Next
	//第一个节点的next置为nil，变为尾节点
	node.Next.Next = nil
	for cur != nil {
		next := cur.Next //保存后续节点
		cur.Next = node.Next 
		node.Next = cur  //插入到头节点后边
		cur = next
	}
}


// 移除有序链表重复元素：快慢指针
func rmDuplicateList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断开与后⾯重复元素的连接
	slow.Next = nil
	return head
}

//k个一组反转链表，剩余链表长度不足k时，无需反转
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a,b := head,head
	for i := 0; i < k; i++ {
		// 不⾜ k 个，不需要反转，base case
		if b == nil {
			return head
		} 
		b = b.Next
	}
	newHead := reverseB(a, b)
	//连接下一段链表的头节点
	a.Next = reverseKGroup(b, k)
	return newHead
}

//翻转区间[a,b)
func reverseB(head *ListNode, b *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
    var pre *ListNode = nil
    for cur != b {
		//直接翻转链表，next指向前置节点
        // pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要

		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
    }
    return pre
}

func TestReverseK(t *testing.T) {
	head := &ListNode{0,nil}
	node := head
	for i := 0; i < 10; i++ {
		node.Next = &ListNode{2 * i, nil}
		node = node.Next
	}
	head.Next.Show()
	fmt.Println("---------------")
	res := reverseKGroup(head.Next, 4)
	res.Show()
}