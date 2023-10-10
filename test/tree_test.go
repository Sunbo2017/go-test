package test

import (
	"fmt"
	"math"
	"testing"
)

type treeNode struct {
	value       string
	left, right *treeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//                                              root
//                           left-1                              right-2
//                 left-3            right-4            left-5             right-6
//             left-7  right-8   left-9  right-10  left-11  right-12   left-13  right-14
func createTree() *treeNode {
	root := &treeNode{"root", nil, nil}

	addLeftNode(1, root)
	addRightNode(2, root)

	addLeftNode(3, root.left)
	addRightNode(4, root.left)
	addLeftNode(5, root.right)
	addRightNode(6, root.right)

	addLeftNode(7, root.left.left)
	addRightNode(8, root.left.left)
	addLeftNode(9, root.left.right)
	addRightNode(10, root.left.right)
	addLeftNode(11, root.right.left)
	addRightNode(12, root.right.left)
	addLeftNode(13, root.right.right)
	addRightNode(14, root.right.right)

	return root
}

func createTreeNode() *TreeNode {
	root := &TreeNode{0, nil, nil}

	addLeft(1, root)
	addRight(2, root)

	addLeft(3, root.Left)
	addRight(4, root.Left)
	addLeft(5, root.Right)
	addRight(6, root.Right)

	addLeft(7, root.Left.Left)
	addRight(8, root.Left.Left)
	addLeft(9, root.Left.Right)
	addRight(10, root.Left.Right)
	addLeft(11, root.Right.Left)
	addRight(12, root.Right.Left)
	addLeft(13, root.Right.Right)
	addRight(14, root.Right.Right)

	return root
}

func addLeftNode(value int, node *treeNode) {
	node.left = &treeNode{value: fmt.Sprintf("left-%d", value)}
}

func addRightNode(value int, node *treeNode) {
	node.right = &treeNode{value: fmt.Sprintf("right-%d", value)}
}

func addLeft(value int, node *TreeNode) {
	node.Left = &TreeNode{Val: value}
}

func addRight(value int, node *TreeNode) {
	node.Right = &TreeNode{Val: value}
}

// 先序遍历
func (node *treeNode) firstTraverse() {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	node.left.firstTraverse()
	node.right.firstTraverse()
}

// 中序遍历
func (node *treeNode) middleTraverse() {
	if node == nil {
		return
	}
	node.left.middleTraverse()
	fmt.Println(node.value)
	node.right.middleTraverse()
}

// 后序遍历
func (node *treeNode) lastTraverse() {
	if node == nil {
		return
	}
	node.left.lastTraverse()
	node.right.lastTraverse()
	fmt.Println(node.value)
}

var result [][]string

// 层序遍历
func (node *treeNode) levelTraverse() [][]string {
	if node == nil {
		return result
	}
	dfsHelper(node, 0)
	return result
}

func dfsHelper(node *treeNode, level int) {
	if node == nil {
		return
	}
	if len(result) < level+1 {
		result = append(result, make([]string, 0))
	}
	result[level] = append(result[level], node.value)
	dfsHelper(node.left, level+1)
	dfsHelper(node.right, level+1)
}

//bfs实现层序遍历
func levelOrder(root *TreeNode) [][]int {
	// write code here
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		l := len(queue)
		//此处遍历每一层
		for i := l; i > 0; i-- {
			node := queue[0]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		res = append(res, level)
	}
	return res
}

//之字形层序遍历二叉树：第一层从左向右，下一层从右向左，一直这样交替
func levelOrder之(root *TreeNode) [][]int {
	// write code here
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	f := 0
	for len(queue) > 0 {
		f++
		level := []int{}
		l := len(queue)
		//此处遍历每一层
		for i := l; i > 0; i-- {
			node := queue[0]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
		}
		if f%2 == 0 {
			//将level逆序
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		res = append(res, level)
	}
	return res
}

func TestLevelOrder(t *testing.T) {
	node := createTreeNode()
	res := levelOrder(node)
	t.Log(res)
	t.Log("---------")
	res = levelOrder之(node)
	t.Log(res)
}

func TestTreeNode(t *testing.T) {
	node := createTree()
	// root 1 3 7 8 4 9 10 2 5 11 12 6 13 14
	node.firstTraverse()
	fmt.Println("---------------------------------------")
	// 7 3 8 1 9 4 10 root 11 5 12 2 13 6 14
	node.middleTraverse()
	fmt.Println("---------------------------------------")
	// 7 8 3 9 10 4 1 11 12 5 13 14 6 2 root
	node.lastTraverse()
}

// depth: current depth
// max: max depth
func dfsCreate(p *treeNode, depth, max int) {
	if depth < max {
		left := &treeNode{value: fmt.Sprintf("%d", 2*depth)}
		right := &treeNode{value: fmt.Sprintf("%d", 4*depth)}
		p.left = left
		p.right = right
		dfsCreate(p.left, depth+1, max)
		dfsCreate(p.right, depth+1, max)
	}
}

func TestDfsCreate(t *testing.T) {
	root := &treeNode{"root", nil, nil}
	dfsCreate(root, 1, 3)
	fmt.Println(root.levelTraverse())
}

//求二叉树最大深度
func maxTreeDepth(root *treeNode) int {
	if root == nil {
		return 0
	}
	//返回子树深度+1
	return Max(maxTreeDepth(root.left), maxTreeDepth(root.right)) + 1
}

func TestMaxDepth(t *testing.T) {
	tree := createTree()
	depth := maxTreeDepth(tree)
	t.Log(depth)
}

//返回二叉树的左侧视图
func leftSideView(root *TreeNode) []int {
	// write code here
	leftView(root)
	return leftViewRes
}

var leftViewRes []int

func leftView(root *TreeNode) {
	if root == nil {
		return
	}

	leftViewRes = append(leftViewRes, root.Val)

	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Left != nil {
		leftView(root.Left)
	} else if root.Right != nil {
		leftView(root.Right)
	}
}

// 翻转二叉树 递归
func mirrorBTree1(root *treeNode) {

	if root.left == nil && root.right == nil {
		return
	}

	root.left, root.right = root.right, root.left

	mirrorBTree1(root.left)
	mirrorBTree1(root.right)
}

//判断两个二叉树是否相同
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

//判断二叉树是否轴对称
func isSymmetric(root *TreeNode) bool {
	// write code here
	if root == nil {
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(root1, root2 *TreeNode) bool {
	if root1 == nil {
		return root2 == nil
	}
	if root2 == nil {
		return root1 == nil
	}
	if root1.Val != root2.Val {
		return false
	}
	return isSymmetricTree(root1.Left, root2.Right) && isSymmetricTree(root1.Right, root2.Left)
}

// 翻转二叉树 遍历
func mirrorBTree2(root *treeNode) {
	stack := NewStack()
	stack.Push(root)

	for stack.Len() > 0 {
		node := stack.Pop().(*treeNode)
		node.left, node.right = node.right, node.left

		if node.left != nil {
			stack.Push(node.left)
		}
		if node.right != nil {
			stack.Push(node.right)
		}
	}

}

// 使用栈实现深度优先
func dfsWithStack(root *treeNode) {
	if root == nil {
		return
	}

	stack := NewStack()
	stack.Push(root)
	for stack.Len() > 0 {
		node := stack.Pop().(*treeNode)
		// 处理当前节点
		fmt.Println(node.value)

		// 先压入右节点
		if node.right != nil {
			stack.Push(node.right)
		}

		// 再压入左节点
		if node.left != nil {
			stack.Push(node.left)
		}
	}
}

//队列实现广度优先
func bfsWithQueue(root *treeNode) {
	if root == nil {
		return
	}

	queue := []*treeNode{root}

	for len(queue) > 0 {
		node := queue[0]

		fmt.Println(node.value)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
		queue = queue[1:]
	}
}

func TestDfsWithStack(t *testing.T) {
	node := createTree()
	// root 1 3 7 8 4 9 10 2 5 11 12 6 13 14
	dfsWithStack(node)
}

func TestBfsWithQueue(t *testing.T) {
	node := createTree()
	// root 1 2 3 4 5 6 7 8 9 10 11 12 13 14
	bfsWithQueue(node)

	// mirrorBTree1(node)
	mirrorBTree2(node)

	// root 2 1 6 5 4 3 14 13 12 11 10 9 8 7
	bfsWithQueue(node)
}

// 判断二叉树是不是二叉查找树
// 中序遍历思想：如果是二叉查找树，中序遍历结果是有序的
var array = []string{}

func judgeSearchTree(root *treeNode) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		return
	}
	judgeSearchTree(root.left)
	array = append(array, root.value)
	judgeSearchTree(root.right)
}

// 直接递归判断,当前节点的值是左子树的最大值，同时是右子树的最小值
func judgeSearchTree1(root *treeNode, maxVal, minVal string) bool {
	if root == nil {
		return true
	}
	if root.left == nil && root.right == nil {
		return true
	}
	if root.value < minVal || root.value > maxVal {
		return false
	}
	if !judgeSearchTree1(root.left, root.value, minVal) {
		return false
	}
	if !judgeSearchTree1(root.right, maxVal, root.value) {
		return false
	}
	return true
}

//中序遍历
//提前记录前一个节点的值
var prev string = "min"

func judgeSearchTree2(root *treeNode) bool {
	if root == nil {
		return true
	}

	if !judgeSearchTree2(root.left) {
		return false
	}

	if root.value < prev {
		return false
	}
	prev = root.value

	return judgeSearchTree2(root.right)
}

func TestJudgeSearchTree(t *testing.T) {
	node := createTree()
	judgeSearchTree(node)
	//判断array是否是正序即可，
	//因为我的treeNode结构value为string，省略判断步骤
	fmt.Println(array)
}

//在BST中判断一个数是否存在(我的treeNode结构value是字符串，所以选用二分法不太合适，只需理解算法逻辑即可)
func judgeExistInTree(tree *treeNode, target string) bool {
	if tree == nil {
		return false
	}
	if tree.value == target {
		return true
	}
	// 参考二分查找算法，提升查找效率
	if target > tree.value {
		return judgeExistInTree(tree.right, target)
	} else {
		return judgeExistInTree(tree.left, target)
	}
}

func TestJudgeSearchTree1(t *testing.T) {
	node := createTree()
	//因为我的treeNode结构value为string，暂写为如下形式
	result := judgeSearchTree1(node, "100", "0")
	fmt.Println(result)
}

// LeetCode 99: ⼆叉搜索树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树
func recoverTree(root *TreeNode) {
	var prev, target1, target2 *TreeNode
	_, target1, target2 = inOrderTraverse(root, prev, target1, target2)
	if target1 != nil && target2 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}
}
func inOrderTraverse(root, prev, target1, target2 *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if root == nil {
		return prev, target1, target2
	}
	prev, target1, target2 = inOrderTraverse(root.Left, prev, target1, target2)
	if prev != nil && prev.Val > root.Val {
		if target1 == nil {
			target1 = prev
		}
		target2 = root
	}
	prev = root
	prev, target1, target2 = inOrderTraverse(root.Right, prev, target1, target2)
	return prev, target1, target2
}

// LeetCode 124：给出⼀个⼆叉树，要求找⼀条路径使得路径的和是最⼤的
//本题中，路径被定义为⼀条从树中任意节点出发，达到任意节点的序列。该路径⾄少包含⼀个节点，且不⼀定经过根节点
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	getPathSum(root, &max)
	return max
}
func getPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MinInt32
	}
	left := getPathSum(root.Left, maxSum)
	right := getPathSum(root.Right, maxSum)
	currMax := Max(Max(left+root.Val, right+root.Val), root.Val)
	*maxSum = Max(*maxSum, Max(currMax, left+right+root.Val))
	return currMax
}

//给定一个数组，判断该数组是不是二叉搜索树的后续遍历结果
func verifyPostorder(postorder []int) bool {

	n := len(postorder)

	if n == 0 {
		return true
	}

	// 最后一个元素为根节点
	root := postorder[n-1]

	// 在二叉搜索树中左子树节点的值小于根节点的值，从当前区域找到第一个大于根节点的，说明后续区域数值都在右子树中
	i := 0
	for ; i < n-1; i++ {
		if postorder[i] > root {
			break
		}
	}

	// 在二叉搜索树中右子树节点的值大于根节点的值
	j := i
	for ; j < n-1; j++ {
		if postorder[j] < root {
			return false
		}
	}

	// 分别判断左右子树是否满足二叉搜索树
	left := true
	if i > 0 {
		left = verifyPostorder(postorder[:i])
	}

	right := true
	if i < n-1 {
		right = verifyPostorder(postorder[i : n-1])
	}

	return left && right
}
