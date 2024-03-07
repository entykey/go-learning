package main

import (
	"fmt"
	"time"
)

func sumArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	nums := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		nums[i] = i + 1
	}

	startTime := time.Now()
	result := sumArray(nums)
	duration := time.Since(startTime)

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Go Execution time: %v\n", duration)
}
