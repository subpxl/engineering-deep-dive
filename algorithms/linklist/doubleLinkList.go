package main

import "fmt"

func DoubleLinkListInit() {
	dll := &DoubleLinkList{}
	dll.append(100)
	dll.append(200)
	dll.append(300)
	dll.append(400)
	dll.TraverseForward()
	fmt.Println("------------------reverse---------------")
	dll.TraverseBackward()
}

type DoubleNode struct {
	Data int
	Next *DoubleNode
	Prev *DoubleNode
}

type DoubleLinkList struct {
	Data int
	Head *DoubleNode
	Tail *DoubleNode
}

func (l *DoubleLinkList) prepend(data int) {
	temp := &DoubleNode{Data: data, Next: nil, Prev: nil}
	if l.Head == nil {
		l.Head = temp
		l.Tail = temp
	} else {
		temp.Next = l.Head
		l.Head.Prev = temp
		l.Head = temp
	}

}

func (l *DoubleLinkList) append(data int) {
	temp := &DoubleNode{Data: data, Next: nil, Prev: nil}
	if l.Head == nil {
		l.Head = temp
		l.Tail = temp
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		temp.Prev = current
		current.Next = temp
		l.Tail = temp
	}
}

func (l *DoubleLinkList) TraverseForward() {
	if l.Head == nil {
		fmt.Println("no lisst")
	} else {
		current := l.Head
		for current != nil {

			fmt.Printf("value = %v, prev = %v, next = %v\n", current.Data, current.Prev, current.Next)
			current = current.Next
		}
	}
}

func (l *DoubleLinkList) TraverseBackward() {
	if l.Head == nil {
		fmt.Println("no list")
	} else {
		current := l.Tail
		for current != nil {
			fmt.Println(current.Data, current.Next, current.Prev)
			current = current.Prev
		}
	}
}
