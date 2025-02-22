package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {

}

func (L *LinkedList) insertdata(data int) {

	// Creating New Node required 2 things data and pointer to next node
	newNode := &Node{data: data, next: nil}

	// If LinkedList is empty create a head node first
	if L.head == nil {

		// Head node will have data and pointer to next node
		L.head = newNode
		return
	}

	// To insert new node we have to find a node which will next point to nil
	currentNode := L.head
	for currentNode.next != nil {
		// You become next node where current node is pointing at and traverse through the linked list
		currentNode = currentNode.next

	}

	// Now you reach a point where current node is not pointing to the next node means next node is nil/empty
	currentNode = newNode

}

func (l *LinkedList) insertInFrontOfHead(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if l.head == nil {

		l.head = newNode
		return
	}

	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) insertNewNode(afterValue, data int) {

	newNode := &Node{
		data: data,
		next: nil,
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head

	for currentNode.next != nil {

		if currentNode.data == afterValue {
			newNode.next = currentNode.next
			currentNode.next = newNode
			return
		}
		currentNode = currentNode.next
	}

	fmt.Printf("Sorry we cannot find the value %d in the linkedlist you were looking for", afterValue)

}
