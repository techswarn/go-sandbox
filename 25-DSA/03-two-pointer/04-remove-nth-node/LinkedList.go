package main

import (
	"fmt"
)

type Linkedlist struct {
	Head *Linkedlist
}

func NewLinkedList(values []int) *Linkedlist {
	ll := &Linkedlist{}
	if len(values) > 0 {
		//Create the linked list by calling the method
		CreateLinkedList(values)
	} 

	return ll
}

// func (ll *Linkedlist) CreateLinkedList(values []int) {
	
// }