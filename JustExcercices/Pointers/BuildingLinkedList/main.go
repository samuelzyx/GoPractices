package main

import "fmt"

// Implement a basic linked list structure using pointers.
//Include functions for adding, removing, and traversing the list.

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Add(data interface{}) {
	newNode := &Node{Data: data}

	// If the list is empty, set the Head to the new node
	if list.Head == nil {
		list.Head = newNode
		return
	}

	// Traverse to the last node
	currentNode := list.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	// Set the last node's Next pointer to the new node
	currentNode.Next = newNode
}

func (list *LinkedList) Remove(data interface{}) bool {
	if list.Head == nil {
		return false
	}
	currentNode := list.Head
	// Check if the head node needs removal
	if currentNode.Data == data {
		list.Head = list.Head.Next
		return true
	}
	// Traverse and search for the node
	for currentNode.Next != nil && currentNode.Next.Data != data {
		currentNode = currentNode.Next
	}
	// If the node is found, update pointers
	if currentNode.Next != nil {
		currentNode.Next = currentNode.Next.Next
		return true
	}
	return false
}

func (list *LinkedList) PrintList() {
	currentNode := list.Head
	for currentNode != nil {
		fmt.Print(currentNode.Data, "->")
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}

func main() {
	//Create new linked list
	linkedList := &LinkedList{}

	//Add some nodes
	linkedList.Add("Apple")
	linkedList.Add("Banana")
	linkedList.Add("Orange")

	fmt.Println("List after adding elements:")
	linkedList.PrintList()

	//remove some element
	removed := linkedList.Remove("Banana")

	if removed {
		fmt.Println("List after remove Banana")
		linkedList.PrintList()
	} else {
		fmt.Println("Couldn't remove banana from the list")
	}
}
