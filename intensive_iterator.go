package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var sum int64
	for i := int64(1); i <= 1000000; i++ {
		sum += i
	}

	duration := time.Since(start)
	fmt.Println("Go Result:", sum)
	fmt.Println("Go Execution time:", duration.Microseconds(), "microseconds.")
}
