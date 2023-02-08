package main

import (
	"fmt"
)

type Any interface{}
type Node struct {
	Value Any
	Next  *Node
}

func add(node *Node, Value Any) *Node {
	node.Value = Value
	node.Next = &Node{}
	return node.Next

}

func insert(head *Node, index int, value Any) {
	var fast *Node
	var last *Node
	var node *Node = &Node{Value: value}

	if index == 0 {
		node.Value = head.Value
		head.Value = value
		last = head.Next
		head.Next = node
		node.Next = last
	} else {
		fast = head.Next
		var count int = 1

		for head.Value != nil {
			if count < index-1 {
				count = count + 1
			} else {
				break
			}
			fast = fast.Next
		}
		fmt.Println("value", fast.Value)
		if fast.Next.Value == nil {
			panic("You cross the size limit of LinkedList. Please use add function for adding values.")
		}

		last = fast.Next
		fast.Next = node
		node.Next = last
	}
}

func print(node *Node) {
	for node.Value != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

func main() {
	head := &Node{Value: 10}
	head.Next = &Node{}
	node := head.Next
	node = add(node, 1)
	node = add(node, 17)
	node = add(node, 3.06)
	node = add(node, "hello somnath")
	insert(head, 4, 44)
	node = add(node, 3.9)

	// Traverse the linked list and print its values
	print(head)

}
