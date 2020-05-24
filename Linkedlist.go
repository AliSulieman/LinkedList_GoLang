/*
author: Ali Sulieman
Date : May 24th, 2020
LinkedList
*/
package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Data     interface{}
	NextNode *Node
}
type Linkedlist struct {
	Head *Node
	Tail *Node
	Size int
}

// inserts an element at the end of the list
func (list *Linkedlist) Insert(element interface{}) {
	node := Node{Data: element}
	if list.Head == nil {
		list.Head = &node
		list.Tail = list.Head
	} else {
		list.Tail.NextNode = &node
		list.Tail = &node
	}
	list.Size++
}

// this function deletes elements from the list
func (list *Linkedlist) delete(element interface{}) {
	previousNode := list.Head
	curentNode := list.Head.NextNode
	if reflect.DeepEqual(list.Head.Data, element) {
		list.Head = list.Head.NextNode
		list.Size--
		return
	}

	for curentNode != nil {
		if reflect.DeepEqual(curentNode.Data, element) {
			previousNode.NextNode = curentNode.NextNode
			curentNode.NextNode = nil
			list.Size--
			return
		}
		previousNode = curentNode
		curentNode = curentNode.NextNode
	}

}
func (list *Linkedlist) getSize() int {
	return list.Size
}
func (list Linkedlist) String() string {

	current := list.Head
	val := fmt.Sprint(current.Data)
	current = current.NextNode
	val += "\n"
	for current != nil {
		val += fmt.Sprint(current.Data)
		val += "\n"
		current = current.NextNode
	}
	return val
}

func main() {
	ll := Linkedlist{}
	ll.Insert(5)
	ll.Insert(7)
	ll.Insert(8)
	ll.delete(7)
	ll.delete(5)
	fmt.Println(ll)
}
