package main

import "fmt"

type linkedListNode struct {
	data int
	next *linkedListNode
}

type Stack struct {
	top *linkedListNode
}

func IsStackEmpty(st *Stack) bool {
	return st.top == nil
}

func Push(st *Stack, element int) {
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

func Pop(st *Stack) int {
	if IsStackEmpty(st) {
		return -1
	} else {
		temp := st.top
		st.top = st.top.next
		temp.next = nil
		return temp.data
	}
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

var visited = [8]int{}
var list = graph()

func DFS(i int) {
	if visited[i] ==0 {
		fmt.Println(i)
		visited[i] = 1
		for v:= 1; v < 8; v++ {
			if list[i][v] == 1 {
				DFS(v)
			}
		}
	}
}

func main() {
	DFS(1)
}
