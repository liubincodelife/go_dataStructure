package main

import "sync"

//数组队列，先进先出
type ArrayQueue struct {
	array []string   //底层切片
	size  int        //队列元素数量
	lock  sync.Mutex //互斥锁
}

//入队 时间复杂度：O(n)
func (q *ArrayQueue) Add(val string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.array = append(q.array, val)
	q.size += 1
}

//出队 时间复杂度：O(n)
func (q *ArrayQueue) Remove() string {
	q.lock.Lock()
	defer q.lock.Unlock()

	//判断队列元素是否为空
	if q.size == 0 {
		panic("ArrayQueue is empty")
	}

	//取队列最前面的元素
	val := q.array[0]
	newArray := make([]string, q.size-1, q.size-1)
	for i := 1; i < q.size; i++ {
		newArray[i-1] = q.array[i]
	}
	q.array = newArray
	q.size -= 1
	return val
}

func (q *ArrayQueue) Size() int {
	return q.size
}
