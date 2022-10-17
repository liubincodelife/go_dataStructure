package main

import "fmt"

func main() {
	//数组队列测试
	fmt.Println("ArrayQueue............")
	arrayQueue := new(ArrayQueue)
	arrayQueue.Add("cat")
	arrayQueue.Add("dog")
	arrayQueue.Add("fish")

	fmt.Println("size:", arrayQueue.Size())
	fmt.Println("remove:", arrayQueue.Remove())
	fmt.Println("remove:", arrayQueue.Remove())
	fmt.Println("size:", arrayQueue.Size())

	//数组队列测试
	fmt.Println("LinkQueue............")
	linkQueue := new(LinkQueue)
	linkQueue.Add("cat")
	linkQueue.Add("dog")
	linkQueue.Add("fish")

	fmt.Println("size:", linkQueue.Size())
	fmt.Println("remove:", linkQueue.Remove())
	fmt.Println("remove:", linkQueue.Remove())
	fmt.Println("size:", linkQueue.Size())
}
