package main

import "fmt"

type ListNode struct {
	Value int32
	Next  *ListNode
}

func main() {
	node := new(ListNode)
	node.Value = 1

	node1 := new(ListNode)
	node1.Value = 3
	node.Next = node1

	node2 := new(ListNode)
	node2.Value = 4
	node1.Next = node2

	curNode := node
	for {
		if curNode != nil {
			fmt.Println(curNode.Value)
			curNode = curNode.Next
		}
		if curNode == nil {
			break
		}
	}
}
