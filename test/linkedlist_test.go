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
	head := &ListNode{0, nil}
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
		node.Next = cur //插入到头节点后边
		cur = next
	}
}

//直接反转链表：当前节点的next指向前一个节点
func reverseLinkedlist1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	cur := head
	var pre *ListNode = nil
	for cur != nil {

		//直接反转
		// pre, cur, cur.Next = cur, cur.Next, pre

		//暂存next指针
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
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
	a, b := head, head
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

func TestReverseList(t *testing.T) {
	head := &ListNode{0, nil}
	node := head
	for i := 0; i < 10; i++ {
		node.Next = &ListNode{2 * i, nil}
		node = node.Next
	}
	head.Next.Show()
	fmt.Println("---------------")
	res := reverseLinkedlist1(head.Next)
	res.Show()
}

func TestReverseK(t *testing.T) {
	head := &ListNode{0, nil}
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

//判断链表是否是回文链表
//如果可以使用额外存储结构的话，可以遍历链表把每个节点的元素都存入一个数组，然后判断数组是否是回文数组就可以
//如果不使用额外数组空间，可以使用链表的后续遍历，模仿双指针操作来判断，代码如下
var left *ListNode

func judgeListBack(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool {
	//递归结束条件
	if right == nil {
		return true
	}

	res := traverse(right.Next)
	//后序遍历代码:后续遍历可以视为栈操作，会最先取到最后节点的值
	res = res && (left.Val == right.Val)
	left = left.Next
	return res
}

//使用快慢指针技巧可以在完全不使用额外空间的情况下实现判断，
//快慢指针找到链表中点，然后反转后半段链表和原链表进行比较，代码如下：
func judgeListBack1(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//遍历结束后，slow位于链表中点
	if fast != nil {
		//fast指针不为nil,说明链表长度为奇数，slow再向前走一步
		slow = slow.Next
	}
	//从slow节点开始反转链表
	right := reverseLinkedlist1(slow)
	left := head
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

//反转部分链表
//将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)，空间复杂度 O(1)
//例如：给出的链表为 1→2→3→4→5→NULL, m=2,n=4,返回 1→4→3→2→5→NULL
//思路：先找到第m个节点cur和前置节点pre，循环将cur后续节点插入到pre和cur之间
func reverseListSection(head *ListNode, m, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	//引入头节点，规避m=1时的边界问题
	h := &ListNode{0, head}
	cur, pre := h, h
	//获取第m个节点和前置节点
	for i := 1; i <= m-1; i++ {
		pre = pre.Next
	}
	cur = pre.Next

	for i := m; i < n; i++ {
		//暂存下一个节点
		nxt := cur.Next
		cur.Next = nxt.Next
		//将后续节点插入到cur之前
		nxt.Next = pre.Next
		pre.Next = nxt
	}
	return h.Next
}

func TestReverseSection(t *testing.T) {
	head := &ListNode{0, nil}
	node := head
	for i := 1; i < 10; i++ {
		node.Next = &ListNode{2 * i, nil}
		node = node.Next
	}
	head.Show()
	fmt.Println("---------------")
	res := reverseListSection(head, 1, 5)
	res.Show()
}

//输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。
//双指针：让N1和N2一起遍历，当N1先走完链表1的尽头（为null）的时候，则从链表2的头节点继续遍历；
//同样，如果N2先走完了链表2的尽头，则从链表1的头节点继续遍历，也就是说，N1和N2都会遍历链表1和链表2。
//因为两个指针，同样的速度，走完同样长度（链表1+链表2），不管两条链表有无相同节点，都能够同时到达终点。
//如果有公共节点，n1和n2会同时走到公共节点处
func FindFirstCommonNode(p1, p2 *ListNode) *ListNode {
	l1, l2 := p1, p2
	for l1 != l2 {
		if l1 == nil {
			l1 = p2
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = p1
		} else {
			l2 = l2.Next
		}
	}
	return l1
}

//链表奇偶重排
//input:1->4->6->5->8->7->nil
//output:1->6->8->4->5->7->nil
//step 1：判断空链表的情况，如果链表为空，不用重排。
//step 2：使用双指针odd和even分别遍历奇数节点和偶数节点，并给偶数节点链表一个头。
//step 3：上述过程，每次遍历两个节点，且even在后面，因此每轮循环用even检查后两个元素是否为NULL，如果不为再进入循环进行上述连接过程。
//step 4：将偶数节点头接在奇数最后一个节点后，再返回头部
func oddEvenList(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	//提取第二个节点
	even := node.Next
	//记录第一个节点
	odd := node
	//为偶数节点设置头节点
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return node
}

func TestOddEvenList(t *testing.T) {
	head := &ListNode{0, nil}
	node := head
	for i := 1; i < 10; i++ {
		node.Next = &ListNode{2 * i, nil}
		node = node.Next
	}
	head.Show()
	fmt.Println("---------------")
	res := oddEvenList(head)
	res.Show()
}
