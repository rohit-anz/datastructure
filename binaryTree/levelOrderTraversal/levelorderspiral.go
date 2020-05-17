package main

import "fmt"

type Node struct{
	data int
	leftChild *Node
	rightChild *Node
}

func newNode(data int) *Node {
	node := &Node{
		data:       data,
		leftChild:  nil,
		rightChild: nil,
	}
	return node
}

func levelOrderSpiral(node *Node) {
	stack1 := Stack{}
	stack2 := Stack{}
	flag := false
	stack1.Push(node)

	for !stack1.IsStackEmpty() || !stack2.IsStackEmpty() {
		if !flag {
			for !stack1.IsStackEmpty() {
				element := stack1.Pop()
				fmt.Println(element.data)
				stack2.Push(element.rightChild)
				stack2.Push(element.leftChild)
			}
		} else {
			for !stack2.IsStackEmpty() {
				element := stack2.Pop()
				fmt.Println(element.data)
				stack1.Push(element.leftChild)
				stack1.Push(element.rightChild)
			}
		}
		flag = !flag
	}
}

func main() {
	root := newNode(1)
	root.leftChild = newNode(2);
	root.rightChild = newNode(3);
	root.leftChild.leftChild = newNode(7);
	root.leftChild.rightChild = newNode(6);
	root.rightChild.leftChild = newNode(5);
	root.rightChild.rightChild = newNode(4);
	levelOrderSpiral(root)
	//preorder(root)
}
