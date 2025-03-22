package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

//Insert Node at the head
func(l *LinkedList) InsertNodeAtHead(node *Node) {
	if l.head = nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

//Create linked list
// func (l *LinkedList) CreateLinkedList(lst []int) {
// 	for i, 
// }
