package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func NewLinkedListNode(data int, next *Node) *LinkedList {
	node := new(Node)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *LinkedList {
	node := new(Node)
	node.data = data
	return node
}