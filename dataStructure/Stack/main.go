package main

import (
	"fmt"
)

func main() {
	//数组栈测试
	fmt.Println("ArrayStack............")
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())

	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())

	fmt.Println("LinkStack............")
	linkStack := new(LinkStack)
	linkStack.Pop()
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())

	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
}
