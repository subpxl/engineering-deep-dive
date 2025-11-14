package main

import (
	"fmt"
)

func main() {

	ar1 := []int{99, 5556, 567, 86, 22, 64, 787, 98, 234, 45}
	// insertionSort(ar1)

	// mergesort
	// sorted := mergeSort(ar1)
	// fmt.Println(sorted)

	soretd := quicksortStart(ar1)
	fmt.Println(soretd)
}

func insertionSort(a []int) {

	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	fmt.Println(a)
}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	right := a[:len(a)/2]
	left := a[len(a)/2:]

	return merge(mergeSort(left), mergeSort(right))

}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	// sort here
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	// merge here
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

// partition
func partition(a []int, low, high int) ([]int, int) {
	pivot := a[high]
	i := low

	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return a, i
}

// quicksort

func quicksort(a []int, low, high int) []int {
	if low < high {
		var p int
		a, p = partition(a, low, high)
		a = quicksort(a, low, p-1)
		a = quicksort(a, high, p+1)
	}
	return a
}

// starter

func quicksortStart(a []int) []int {
	return quicksort(a, 0, len(a)-1)
}
