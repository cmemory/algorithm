package datastructure

// 单链表节点
type LNode struct {
	T interface{}
	Next *LNode
}

// 二叉树节点
type BNode struct {
	T interface{}
	Left *BNode
	Right *BNode
}