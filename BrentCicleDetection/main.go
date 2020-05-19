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
	node5.nextNode = &node2

	// main program
	n := &node1
	power, isCycle := n.getPowerOfTwoValue()
	if isCycle {
		nodePowerOfTwo := n.getNodePowerOfTwo(power)
		id := n.getIDCicleBegin(nodePowerOfTwo)
		fmt.Println(id)
	} else {
		fmt.Println("Cycle did not detect.")
	}
}

func (n *Node) getPowerOfTwoValue() (int, bool) {
	power := 1
	lam := 1
	tortoise := n
	hare := n.nextNode
	for tortoise != hare {
		if power == lam {
			tortoise = hare
			power *= 2
			lam = 0
		}
		hare = hare.nextNode
		lam++
		if hare == nil {
			return 0, false
		}
	}
	return power, true
}

func (n *Node) getNodePowerOfTwo(power int) *Node {
	hare := n
	for power != 0 {
		hare = hare.nextNode
		power--
	}
	return hare
}

func (n *Node) getIDCicleBegin(hare *Node) int {
	mu := 0
	t := n
	h := hare
	for t != h {
		t = t.nextNode
		h = h.nextNode
		mu++
	}
	return mu
}
