package main

import "fmt"

// Linked List
type LinkedList struct {
	Head   *Node
	Length int
}

type Node struct {
	Data int
	Next *Node
}

func main() {

	myl := LinkedList{Head: nil}
	myl.Create(1)
	myl.Create(2)
	myl.Create(3)
	myl.InsertN(4, 3)
	myl.show()
	fmt.Println("-------------------------------------deleted-------------")
	// myl.Delete()
	// myl.deleteN(1)
	// myl.deleteFirst()
	// myl.reverse()
	ReverseRecursion(myl.Head)
	myl.show()

}

func (l *LinkedList) Create(data int) {
	elem := &Node{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = elem
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = elem
	}
}

func (l *LinkedList) show() {
	if l.Head == nil {
		fmt.Println("no data")
	} else {
		current := l.Head
		for current != nil {
			fmt.Println(current)
			current = current.Next
		}
	}
}

func (l *LinkedList) InsertN(data int, n int) {
	elem := &Node{Data: data, Next: nil}
	if l.Head == nil {
		fmt.Println("no data")
	} else {
		count := 0
		current := l.Head
		for current != nil {
			count += 1
			if count == n {
				k := current.Next
				current.Next = elem
				elem.Next = k
			}
			current = current.Next
		}
	}
}

func (l *LinkedList) Delete() {
	if l.Head == nil {
		fmt.Println("no data")
	} else {

		current := l.Head
		for current.Next.Next != nil {
			current = current.Next

		}
		current.Next = nil
	}
}

func (l *LinkedList) deleteN(n int) {
	if l.Head == nil {
		fmt.Println("no data")
	} else {
		count := 0
		current := l.Head
		for current != nil {
			count += 1
			prev := current
			current = current.Next
			if count == n {
				prev.Next = current.Next
			}
		}
	}

}

func (l *LinkedList) deleteFirst() {

	if l.Head == nil {
		fmt.Println("no list")
	} else {
		l.Head = l.Head.Next
	}
}

func (l *LinkedList) reverse() {
	if l.Head == nil {
		fmt.Println("no list")
	} else {
		current := l.Head
		var prev *Node
		var next *Node

		for current != nil {
			next = current.Next //first get next location
			current.Next = prev // set next location to prev
			prev = current      //change pointer a -> b to b-> a
			current = next      //move to third pointer
		}
		l.Head = prev //add head pointer to  last block address

	}

}

func ReverseRecursion(node *Node) {
	p := node
	if p == nil {
		return
	}
	ReverseRecursion(p.Next)
	fmt.Println(p.Data)
}
