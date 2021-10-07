package main

import (
	"fmt"
	"gkTime-goBasic/examples/comm_pkg/queue"
	"gkTime-goBasic/examples/comm_pkg/stack"
)

func main() {
	// 初始化一个二叉树
	nodeG := binaryTree{data: "g", left: nil, right: nil}
	nodeF := binaryTree{data: "f", left: &nodeG, right: nil}
	nodeE := binaryTree{data: "e", left: nil, right: nil}
	nodeD := binaryTree{data: "d", left: &nodeE, right: nil}
	nodeX := binaryTree{data: "x", left: nil, right: nil}
	nodeZ := binaryTree{data: "z", left: nil, right: nil}
	nodeY := binaryTree{data: "y", left: nil, right: &nodeZ}
	nodeC := binaryTree{data: "c", left: &nodeX, right: &nodeY}
	nodeB := binaryTree{data: "b", left: &nodeD, right: &nodeF}
	// root节点，后续传值使用指针，更节省空间
	nodeA := binaryTree{data: "a", left: &nodeB, right: &nodeC}
	root := nodeA
	//fmt.Println("前序遍历(递归)：")
	//root.preOrderRecursive(&root)
	//fmt.Println()

	//fmt.Println("后续遍历(递归)：")
	//root.postOrderRecursive(&root)
	//fmt.Println()

	//fmt.Println("中序遍历(递归)：")
	//root.midOrderRecursive(&root)
	//fmt.Println()
	//
	//fmt.Println("中序遍历(栈)：")
	//root.midOrderStack(&root)
	//fmt.Println()

	fmt.Println("深度优先遍历(递归)：")
	root.dfOrderRecursive(&root)
	fmt.Println()

	fmt.Println("深度优先遍历(栈)：")
	root.dfOrderStack(&root)
	fmt.Println()

	//fmt.Println("广度优先遍历（队列）：")
	//root.BfOrder(&root)
	//fmt.Println()
}

type Tree interface {
	Insert(key string) bool
	Search(key string) bool
	// Modify()
	Delete(key string)
}

// 二叉树
type binaryTree struct {
	data string
	//IsRoot bool
	left  *binaryTree
	right *binaryTree
}

// 多叉树
type multiWayTree struct {
}

// preOrderRecursive 二叉数前序遍历-递归版，遍历顺序：跟节点 -> 左子树 -> 右子树
func (t *binaryTree) preOrderRecursive(tree *binaryTree) {
	if tree == nil {
		return
	}
	fmt.Print(tree.data, " ")
	t.preOrderRecursive(tree.left)
	t.preOrderRecursive(tree.right)
}

func (t *binaryTree) preOrderStack(tree *binaryTree) {
	if tree == nil {
		return
	}

	stack := stack.NewSliceStack()
	var elements []string

	for !stack.IsEmpty() || tree != nil {
		if tree != nil {
			elements = append(elements, tree.data)
			stack.Push(tree)
			tree = tree.left
		} else {
			node := (*stack.Pop()).(*binaryTree)
			tree = node.right
		}
	}

	// 此种栈思想也可以实现
	//stack.Push(tree)
	//for !stack.IsEmpty() {
	//	node := (*stack.Pop()).(*binaryTree)
	//	elements = append(elements, node.data)
	//	// 优先右节点入栈
	//	if node.right != nil {
	//		stack.Push(node.right)
	//	}
	//	if node.left != nil {
	//		stack.Push(node.left)
	//	}
	//}

	for _, e := range elements {
		fmt.Print(e, " ")
	}
}

// midOrderRecursive 二叉树中序遍历-递归版，遍历顺序：左子树 -> 跟节点 -> 右子树
func (t *binaryTree) midOrderRecursive(tree *binaryTree) {
	if tree == nil {
		return
	}
	t.midOrderRecursive(tree.left)
	fmt.Print(tree.data, " ")
	t.midOrderRecursive(tree.right)
}

// midOrderStack 二叉树中序遍历-栈版，遍历顺序：左子树 -> 跟节点 -> 右子树
func (t *binaryTree) midOrderStack(tree *binaryTree) {
	if tree == nil {
		return
	}
	stack := stack.NewSliceStack()
	var elements []string
	//curNode := &binaryTree{}
	//stack.Push(tree)

	for !stack.IsEmpty() || tree != nil {
		if tree != nil {
			stack.Push(tree)
			tree = tree.left
		} else {
			node := (*stack.Pop()).(*binaryTree)
			elements = append(elements, node.data)
			tree = node.right
		}
	}

	for _, e := range elements {
		fmt.Print(e, " ")
	}
}

// postOrderRecursive 二叉树后续遍历-递归版，遍历顺序：左子树 -> 右子树 -> 跟节点
func (t *binaryTree) postOrderRecursive(tree *binaryTree) {
	if tree == nil {
		return
	}
	t.postOrderRecursive(tree.left)
	t.postOrderRecursive(tree.right)
	fmt.Print(tree.data, " ")
}

// postOrderStack 二叉树后续遍历-栈版，遍历顺序：左子树 -> 右子树 -> 跟节点
func (t *binaryTree) postOrderStack(tree *binaryTree) {
	if tree == nil {
		return
	}
	stack := stack.NewSliceStack()
	var elements []string
	stack.Push(tree)
	preNode := &binaryTree{} // 辅助节点，用于判断节点是否输出
	for !stack.IsEmpty() {
		node := (*stack.GetFirstItem()).(*binaryTree)
		// 节点为叶子节点或者全部字节点节点均入栈，则可以输出
		if (node.right == nil && node.left == nil) || (preNode != nil && (preNode == node.right || preNode == node.left)) {
			elements = append(elements, node.data)
			preNode = node
			stack.Pop()
		} else {
			// 先右节点入栈
			if node.right != nil {
				stack.Push(node.right)
			}
			// 左节点入栈
			if node.left != nil {
				stack.Push(node.left)
			}
		}
	}
	for _, e := range elements {
		fmt.Println(e, " ")
	}
}

// BfOrder 二叉树广度优先遍历(层次遍历)，遍历顺序：从根节点一层层从左向右遍历
// 广度遍历的运行过程是先进先出的，自然的方法是队列
func (t *binaryTree) BfOrder(tree *binaryTree) {
	if tree == nil {
		return
	}
	queue := queue.NewSliceQueue()
	var elements []string
	// 根节点入队列
	queue.Push(tree)
	for !queue.IsEmpty() {
		node := (*queue.Pop()).(*binaryTree)
		// 可以打印，也可以放在一个切片中，再进行输出
		elements = append(elements, node.data)
		//fmt.Print(node.data, " ")
		// 左节点入队列
		if node.left != nil {
			queue.Push(node.left)
		}
		// 右节点入队列
		if node.right != nil {
			queue.Push(node.right)
		}
	}

	// 循环输出元素
	for _, e := range elements {
		fmt.Print(e, " ")
	}

}

// dfOrderRecursive 二叉树深度优先遍历-递归版，遍历顺序：和前序遍历一样，遍历顺序：跟节点 -> 左子树 -> 右子树
// 深度遍历的运行过程是先进后出的，自然的方法是栈和递归
func (t *binaryTree) dfOrderRecursive(tree *binaryTree) {
	t.preOrderRecursive(tree)
}

// dfOrderStack 二叉树深度优先遍历-栈版
func (t *binaryTree) dfOrderStack(tree *binaryTree) {
	t.preOrderStack(tree)
}
