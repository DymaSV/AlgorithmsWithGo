package main

import (
	"fmt"
)

type Node struct {
	id       int
	nextNode *Node
}

func main() {
	node1 := Node{id: 0, nextNode: nil}
	node2 := Node{id: 1, nextNode: nil}
	node3 := Node{id: 2, nextNode: nil}
	node4 := Node{id: 3, nextNode: nil}
	node5 := Node{id: 4, nextNode: nil}

	node1.nextNode = &node2
	node2.nextNode = &node3
	node3.nextNode = &node4
	node4.nextNode = &node5
	node5.nextNode = &node3

	fmt.Println(node3)
}
