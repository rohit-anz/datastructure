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

func levelOrderLineByLine(root *Node) {
	queue := Queue{}
	queue.enqueue(root)
	queue.enqueue(nil)

	for queue.size() > 1 {
		element := queue.dequeue()
		if element != nil {
			fmt.Printf("%v ",element.data)
			if element.leftChild != nil {
				queue.enqueue(element.leftChild)
			}

			if element.rightChild != nil {
				queue.enqueue(element.rightChild)
			}
		} else {
			queue.enqueue(nil)
			fmt.Printf("\n")
		}
	}
}

func main() {
/*
	root := newNode(1)
	root.leftChild = newNode(2);
	root.rightChild = newNode(3);
	root.leftChild.leftChild = newNode(7);
	root.leftChild.rightChild = newNode(6);
	root.rightChild.leftChild = newNode(5);
	root.rightChild.rightChild = newNode(4);
*/
	root := newNode(1)
	root.leftChild = newNode(2);
	root.rightChild = newNode(3);
	root.leftChild.leftChild = newNode(4);
	root.leftChild.rightChild = newNode(5);
	root.rightChild.rightChild = newNode(6);

	levelOrderLineByLine(root)
}
