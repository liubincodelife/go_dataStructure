package main

import "sync"

type LinkQueue struct {
	root *ListNode  //链表起始节点
	size int        //队列元素数量
	lock sync.Mutex //互斥锁
}

//链表队列，先进先出
type ListNode struct {
	Value string
	Next  *ListNode
}

//入队 时间复杂度：O(n)
func (q *LinkQueue) Add(val string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.root == nil {
		q.root = new(ListNode)
		q.root.Value = val
	} else {
		newNode := new(ListNode)
		newNode.Value = val
		//查找尾节点
		curNode := q.root
		for curNode.Next != nil {
			curNode = curNode.Next
		}
		//将新节点插入链表尾部
		curNode.Next = newNode
	}
	q.size += 1
}

//出队 时间复杂度：O(1)
func (q *LinkQueue) Remove() string {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		panic("LinkQueue is empty")
	}

	topNode := q.root
	val := topNode.Value
	//删除头节点
	q.root = topNode.Next
	q.size -= 1
	return val
}

func (q *LinkQueue) Size() int {
	return q.size
}
