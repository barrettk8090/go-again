package main

import (
	"fmt"
	"time"
)

// Speed test showing the benefits of setting
// capacity of slices ahead of time (if/when possible)

func main() {
	var n int = 1000000
	// empty capacity
	var testSlice = []int{}
	// pre-allocated capacity of 1000000
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
