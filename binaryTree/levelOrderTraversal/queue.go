package main

import "fmt"

type queueNode struct {
	data *Node
	next * queueNode
}

type Queue struct {
	rear *queueNode
	front *queueNode
}

func (q *Queue) enqueue(data *Node)  {
	if data == nil {
		return
	}

	t := &queueNode{
		data: data,
		next: nil,
	}

	if q.rear == nil {
		q.rear = t
		q.front = t
	} else {
		q.rear.next = t
		q.rear = q.rear.next
	}
}

func (q *Queue) isQueueEmpty() bool {
	if q.front == nil {
		return true
	}
	return false
}

func (q *Queue) dequeue() *Node  {
	if q.front == nil {
		fmt.Println("queue is empty")
		return nil
	} else  if q.front == q.rear && q.front != nil {
		element:= q.front.data
		q.front = nil
		q.rear = nil
		return element
	} else {
		element := q.front.data
		q.front = q.front.next
		return element
	}
}
