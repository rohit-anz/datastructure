package main

import "fmt"

type node struct {
	data int
	next * node
}

type queue struct {
	rear *node
	front *node
}

func (q *queue) enqueue(data int)  {
	t := &node{
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

func (q *queue) isQueueEmpty() bool {
	if q.front == nil {
		return true
	}
	return false
}

func (q *queue) dequeue() int  {
	if q.front == nil {
		fmt.Println("queue is empty")
		return -1
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


func main() {
	breadthFirstSearch(1)
}

func graph() [8][8]int {
	return [8][8]int{{0,0,0,0,0,0,0,0},
		{0,0,1,1,1,0,0,0},
		{0,1,0,1,0,0,0,0},
		{0,1,1,0,1,1,0,0},
		{0,1,0,1,0,1,0,0},
		{0,0,0,1,1,0,1,1},
		{0,0,0,0,0,1,0,0},
		{0,0,0,0,0,1,0,0}}
}

func breadthFirstSearch(i int) {
	q := queue{}
	visited := [8]int{}
	list := graph()
	fmt.Println(i)
	q.enqueue(i)
	visited[i] = 1

	for !q.isQueueEmpty() {
		element := q.dequeue()
		for j:= 1; j < 8; j++ {
			if list[element][j] == 1 && visited[j] == 0 {
				fmt.Println(j)
				q.enqueue(j)
				visited[j] = 1
			}
		}
	}
}