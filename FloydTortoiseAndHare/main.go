package main

import (
	"fmt"
)

type Node struct {
	id       int
	nextNode *Node
}

func main() {
	//create LinkedList
	node1 := Node{id: 0, nextNode: nil}
	node2 := Node{id: 1, nextNode: nil}
	node3 := Node{id: 2, nextNode: nil}
	node4 := Node{id: 3, nextNode: nil}
	node5 := Node{id: 4, nextNode: nil}
	node1.nextNode = &node2
	node2.nextNode = &node3
	node3.nextNode = &node4
	node4.nextNode = &node5
	node5.nextNode = nil

	// main program
	n := &node1
	node, isCycle := n.getHareNode(n)
	if isCycle {
		id := n.getCycleID(node)
		fmt.Println(id)
	} else {
		fmt.Println("Cycle did not detect.")
	}
}

func (n *Node) getHareNode(hare *Node) (*Node, bool) {
	if n.nextNode != nil {
		t := n.nextNode
		h := hare.nextNode
		if h != nil && h.nextNode != nil {
			h = h.nextNode
			if t == h {
				return h, true
			}
			return t.getHareNode(h)
		}
	}
	return nil, false
}

func (n *Node) getCycleID(hare *Node) int {
	t := n.nextNode
	h := hare.nextNode
	if t == h {
		return t.id
	}
	return t.getCycleID(h)
}
