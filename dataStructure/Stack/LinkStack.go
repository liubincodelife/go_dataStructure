package main

import "sync"

type LinkStack struct {
	root *ListNode
	size int
	lock sync.Mutex
}

type ListNode struct {
	Next  *ListNode
	Value string
}

//入栈
func (s *LinkStack) Push(val string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	//如果栈顶为空，则增加节点
	if s.root == nil {
		s.root = new(ListNode)
		s.root.Value = val
	} else {
		//新插入的元素放在链表的头部
		preNode := s.root
		newNode := new(ListNode)
		newNode.Value = val

		//原链表头节点链接到新节点后面
		newNode.Next = preNode
		s.root = newNode
	}
	s.size += 1
}

func (s *LinkStack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()

	//判断栈是否为空
	if s.size == 0 {
		panic("stack is empty")
	}
	topNode := s.root
	val := topNode.Value

	s.root = topNode.Next
	//栈中元素大小减1
	s.size -= 1
	return val
}

//获取栈顶元素
func (s *LinkStack) Peek() string {
	//判断栈是否为空
	if s.size == 0 {
		panic("empty link stack")
	}

	val := s.root.Value
	return val
}

//获取栈大小
func (s *LinkStack) Size() int {
	return s.size
}

func (s *LinkStack) IsEmpty() bool {
	return s.size == 0
}
