package main

type Node struct {
	// 自引用只能使用指针，
	// 因为在编译会计算结构体的内存分配，如果结构体自引用会出现死循环，编译也会报错
	left *Node
	right *Node

	// 这个也会报错, NodeNode这个结构体会出现报错
	// nn NodeNode
}


type NodeNode struct {
	node Node
}