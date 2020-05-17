package main

type linkedListNode struct {
	data *Node
	next *linkedListNode
}

type Stack struct {
	top *linkedListNode
}

func (st *Stack) IsStackEmpty() bool {
	return st.top == nil
}

func (st *Stack) Push(element *Node) {
	if element == nil {
		return
	}

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

func (st *Stack) Pop() *Node {
	if st.IsStackEmpty() {
		return nil
	} else {
		temp := st.top
		st.top = st.top.next
		temp.next = nil
		return temp.data
	}
}
