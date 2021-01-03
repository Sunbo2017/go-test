package test

import (
	"fmt"
	"testing"
)

type treeNode struct {
	value       string
	left, right *treeNode
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

func addLeftNode(value int, node *treeNode) {
	node.left = &treeNode{value: fmt.Sprintf("left-%d", value)}
}

func addRightNode(value int, node *treeNode) {
	node.right = &treeNode{value: fmt.Sprintf("right-%d", value)}
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
    if len(result) < level + 1 {
        result = append(result, make([]string, 0))
    }
    result[level] = append(result[level], node.value)
    dfsHelper(node.left, level + 1)
    dfsHelper(node.right, level + 1)
}

func TestTreeNode(t *testing.T){
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
		left := &treeNode{value: fmt.Sprintf("%d", 2 * depth)}
		right := &treeNode{value: fmt.Sprintf("%d", 4 * depth)}
		p.left = left
		p.right = right
		dfsCreate(p.left, depth+1, max)
		dfsCreate(p.right, depth+1, max)
	}
}

func TestDfsCreate(t *testing.T){
	root := &treeNode{"root", nil, nil}
	dfsCreate(root, 1, 3)
	fmt.Println(root.levelTraverse())
}

// 使用栈实现 dfs
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
		if node.right != nil{
			stack.Push(node.right)
		}
		
		// 再压入左节点
		if node.left != nil{
			stack.Push(node.left)
		}
	}
}

func bfsWithQueue(root *treeNode){
	if root == nil{
		return
	}

	queue := []*treeNode{root}

	for len(queue) > 0{
		node := queue[0]

		fmt.Println(node.value)

		if node.left != nil{
			queue = append(queue, node.left)
		}

		if node.right != nil{
			queue = append(queue, node.right)
		}
		queue = queue[1:]
	}
}

func TestDfsWithStack(t *testing.T){
	node := createTree()
	// root 1 3 7 8 4 9 10 2 5 11 12 6 13 14
	dfsWithStack(node)
}

func TestBfsWithQueue(t *testing.T){
	node := createTree()
	// root 1 2 3 4 5 6 7 8 9 10 11 12 13 14
	bfsWithQueue(node)
}