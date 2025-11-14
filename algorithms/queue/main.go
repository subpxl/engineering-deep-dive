package main

import "fmt"

// first in first out
// Queue implementation
type Queue struct {
	Items []int
}

func main() {

	myq := Queue{}
	myq.enqueue(100)
	myq.enqueue(200)
	myq.enqueue(300)
	myq.enqueue(400)
	fmt.Println(myq.Items)
	fmt.Println(myq.dequeue())
	fmt.Println(myq.dequeue())

}

func (q *Queue) enqueue(data int) int {
	q.Items = append(q.Items, data)
	return data

}

func (q *Queue) dequeue() (int, bool) {
	if len(q.Items) == 0 {
		return 0, false
	} else {
		firstItem := q.Items[0]
		q.Items = append(q.Items[1:])
		return firstItem, true
	}
}
