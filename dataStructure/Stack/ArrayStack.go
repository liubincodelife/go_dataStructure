package main

import "sync"

//数组栈，后进先出
type ArrayStack struct {
	array []string   //底层切片
	size  int        //栈的元素数量
	lock  sync.Mutex //互斥锁，并发安全
}

//入栈
func (s *ArrayStack) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	//放入切片中，后进的元素放在数组最后面
	s.array = append(s.array, v)
	//栈中的元素数量+1
	s.size += 1
}

//出栈
func (s *ArrayStack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()

	//如果栈中元素为空
	if s.size == 0 {
		panic("empty stack")
	}

	//取栈顶元素
	val := s.array[s.size-1]
	//直接收缩切片，多余的元素不会释放，会导致占用空间越来越大
	//s.array = s.array[:s.size-1]

	//创建新的数组，时间复杂度为O(n)
	newArray := make([]string, s.size-1, s.size-1)
	for i := 0; i < s.size-1; i++ {
		newArray[i] = s.array[i]
	}
	s.array = newArray
	s.size = s.size - 1
	return val
}

//获取栈顶元素
func (s *ArrayStack) Peek() string {
	//栈中元素为空
	if s.size == 0 {
		panic("empty array stack")
	}

	val := s.array[s.size-1]
	return val
}

//获取栈大小
func (s *ArrayStack) Size() int {
	return s.size
}

//判断栈是否为空
func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}
