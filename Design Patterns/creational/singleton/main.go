package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Couner struct {
	mu    sync.Mutex
	count int64
}

var (
	once     sync.Once
	instance *Couner
)

func GetCounter() *Couner {
	once.Do(func() {
		instance = &Couner{count: 0}
	})
	return instance
}

func (c *Couner) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Couner) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	c1 := GetCounter()
	c2 := GetCounter()
	fmt.Println("same isntance ", c1 == c2)
	for i := 0; i < 5; i++ {
		c1.Increment()
	}
	fmt.Println("c1 count after loop  ", c1.GetCount())
	fmt.Println("c2 count after loop", c2.GetCount())
}
