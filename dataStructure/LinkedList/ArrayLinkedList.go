package main

import "fmt"

type ArrayLinkedList struct {
	Value string
	Next  int32
}

func ArrayLinked() {
	array := [5]ArrayLinkedList{}
	array[0] = ArrayLinkedList{"I", 3}
	array[1] = ArrayLinkedList{"Army", 4}
	array[2] = ArrayLinkedList{"You", 1}
	array[3] = ArrayLinkedList{"Love", 2}
	//-1表示没有后继节点
	array[4] = ArrayLinkedList{"!", -1}

	node := array[0]
	for {
		fmt.Println(node.Value)
		if node.Next == -1 {
			break
		}
		node = array[node.Next]
	}
}
