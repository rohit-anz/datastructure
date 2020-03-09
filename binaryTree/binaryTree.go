package binaryTree

import (
	"fmt"
)

type Node struct{
	data int
	leftChild *Node
	rightChild *Node
}

type linkedListNode struct {
	data *Node
	next *linkedListNode
}

type QueueNode struct {
	data *Node
	next *QueueNode
}

type Stack struct {
	top *linkedListNode
}

type Queue struct {
	rear *QueueNode
	front *QueueNode
}

func IsStackEmpty(st *Stack) bool {
	return st.top == nil
}

func Push(st *Stack, element *Node) {
	temp := linkedListNode{
		data: element,
		next: nil,
	}
	if st.top == nil {
		st.top = &temp
	} else {
		temp.next = st.top
		st.top = &temp
	}
}

func Pop(st *Stack) *Node {
	if IsStackEmpty(st) {
		return nil
	} else {
		temp := st.top
		st.top = st.top.next
		temp.next = nil
		return temp.data
	}
}

func Insert(root *Node, element int) *Node  {
	if root == nil {
		root = &Node{
			leftChild:  nil,
			rightChild: nil,
			data:       element,
		}
		return root
	}

	t:= root
	var r *Node = t

	for; t!= nil; {
		r = t
		if element < t.data {
			t = t.leftChild
		} else {
			t = t.rightChild
		}
	}

	p := Node{
		leftChild:  nil,
		rightChild: nil,
		data:       element,
	}

	if element < r.data {
		r.leftChild = &p
	} else {
		r.rightChild = &p
	}
	return  root
}

func Preorder(root *Node) {
	stack := Stack{top:nil}
	for root != nil || !IsStackEmpty(&stack) {
		if root != nil {
			fmt.Println(root.data)
			Push(&stack, root)
			root = root.leftChild
		} else {
			root = Pop(&stack)
			root = root.rightChild
		}
	}
}

func Inorder(root *Node) {
	stack:= Stack{}

	if root == nil {
		fmt.Println("nothing really exists")
	}else if root.leftChild == nil && root.rightChild == nil {
		fmt.Println(root.data)
	}

	for root != nil || !IsStackEmpty(&stack) {
		if root != nil {
			Push(&stack, root)
			root = root.leftChild
		} else {
			root = Pop(&stack)
			fmt.Println(root.data)
			root = root.rightChild
		}
	}
}

func PostOrder(root *Node) {
	if root != nil {
		PostOrder(root.leftChild)
		PostOrder(root.rightChild)
		fmt.Println(root.data)
	}
}

func insert(queue *Queue, element *Node) {
	node := &QueueNode {
		data: element,
		next: nil,
	}

	if queue.rear == nil && queue.front == nil {
		queue.rear = node
		queue.front = node
	} else {
		queue.rear.next = node
		queue.rear = queue.rear.next
	}
}

func delete(queue *Queue) *QueueNode {
	node:= queue.front
	 queue.front = queue.front.next
	 if queue.front == nil {
	 	queue.rear = nil
	 }
	return node
}

func isQueueEmpty(queue *Queue) bool {
	return queue.rear == nil || queue.front == nil
}

func LevelOrder(root *Node) {
	queue := &Queue{
		rear:  nil,
		front: nil,
	}

	if root != nil {
		insert(queue, root)
		for !isQueueEmpty(queue) {
			node := delete(queue)
			if node.data.leftChild != nil {
				insert(queue, node.data.leftChild)
			}

			if node.data.rightChild != nil {
				insert(queue, node.data.rightChild)
			}
			fmt.Println(node.data.data)
		}
	} else {
		fmt.Println("tree does not exist")
	}
}

func SumOfAllNode(root *Node) int {
	if root != nil {
		x := SumOfAllNode(root.leftChild)
		y := SumOfAllNode(root.rightChild)
		return x + y + root.data
	}
	return 0
}

func CountNumberOfNode(root *Node) int {
	if root != nil {
		x := CountNumberOfNode(root.leftChild)
		y := CountNumberOfNode(root.rightChild)
		return x + y + 1
	}
	return 0
}

func HeightOfTree(root *Node) int {
	if root != nil {
		x := HeightOfTree(root.leftChild) + 1
		y := HeightOfTree(root.rightChild) + 1
		if x > y {
			return x
		} else {
			return y
		}
	}
	return 0
}

func inorderPredecer(root *Node) *Node {
	root = root.leftChild

	for root.rightChild != nil {
		root = root.rightChild
	}

	return root
}

func inorderSuccessor(root *Node) *Node {
	root = root.rightChild

	for root.leftChild != nil {
		root = root.leftChild
	}
	return root
}

func DeleteNode(root *Node, key int) *Node {

	if root == nil {
		return nil
	}

	if root.leftChild == nil && root.rightChild == nil {
		root = nil
		return nil
	}

	if key < root.data {
		root.leftChild = DeleteNode(root.leftChild, key)
	} else if key > root.data {
		root.rightChild = DeleteNode(root.rightChild, key)
	} else {
		if HeightOfTree(root.leftChild) > HeightOfTree(root.rightChild) {
			temp := root
			inorderPredecer := inorderPredecer(temp)
			root.data = inorderPredecer.data
			root.leftChild = DeleteNode(root.leftChild, temp.data)
		} else {
			temp := root
			inorderSuccessor := inorderSuccessor(temp)
			root.data = inorderSuccessor.data
			root.rightChild = DeleteNode(root.rightChild, temp.data)
		}
	}
	return root
}