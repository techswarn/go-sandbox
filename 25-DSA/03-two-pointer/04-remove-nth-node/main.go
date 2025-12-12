package main

import (
	"fmt"
)

//Create a node struct

type Node struct {
	Value int
	next *Node
}

// Create linkedlist struct
type Linkedlist struct {
	Head *Node
}

//function to create a new newLinked list
func NewLinkedList(values []int) *Linkedlist {
	ll := &Linkedlist{}
	if len(values) > 0 {
		//Func to create the linked list
		ll.CreateLinkedList(values)
	}
	return ll
}

func (ll *Linkedlist) CreateLinkedList(values []int) {
	if len(values) == 0 {
		ll.Head = nil 
		return
	}

	ll.Head = &Node{Value: values[0]}
	current := ll.Head
	for _, val := range values[1:]{
		current.next = &Node{Value: val}
		current = current.next
	}

}

func (ll *Linkedlist) DisplayLinkedlist() {
	cur := ll.Head
	for cur != nil {
		fmt.Printf("%d->", cur)
		cur = cur.next
	}
	fmt.Println("None")
}

func main() {
	fmt.Println("Remove nth node from the linked list")
	arr := []int{1,2,3,4,5,6}
	//NewLinkedList(arr)
	List := &Linkedlist{}
	List.CreateLinkedList(arr)
	List.DisplayLinkedlist()
}
