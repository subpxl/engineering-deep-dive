package main

import "fmt"

func main() {

	ans := fibonacci(100)

	fmt.Println(ans)
	fmt.Println(calculations)
}

var calculations = 0

func fibonacci(n int) int {
	cache := make(map[int]int)

	for i := 0; i < n; i++ {
		if i < 3 {
			cache[i] = i
		} else {
			calculations++
			cache[i] = cache[i-1] + cache[i-2]
		}
	}
	return cache[n-1]
}
