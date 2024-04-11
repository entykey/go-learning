// https://dev.to/ghanithan/how-to-squeeze-every-drop-of-cpu-resource-using-concurrent-parallel-programming-in-go-578f
/*
package main

import (
	"fmt"
)

func main() {

	c := make(chan int) // create a channel for the goroutine to send value to the main thread

	go getNum(1, 100, c) // start a goroutine which runs in parallel to the main thread

	fmt.Print("\n")
	for m := range c { // Loop and wait for the value to be received on the channel.
		// Loop breaks when the channel in closed by the sender
		fmt.Print(m, ",")
	}
	fmt.Print("\n\n")

}

// function to generated 100 numbers in sequence and push them in a channel
func getNum(init, end int, c chan int) {

	for i := init; i <= end; i++ {
		c <- i // write value on the channel
	}
	close(c) // close the channel once the task is completed

}
*/

/*
// another example:
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1) // instructing the go runtine to use only one OS thread to observe the concurrency

	c := make(chan int) // create a channel for the goroutine to send value to the main thread

	go getNum(1, 100, c) // start a goroutine which runs in parallel to the main thread

	go printHello(50) // print hello 50 times concurrently

	fmt.Print("\n")
	for m := range c { // Loop and wait for the value to be received on the channel.
		// Loop breaks when the channel in closed by the sender
		fmt.Print(m, ",")
	}

	fmt.Print("\n\n")

}

// function to generated numbers in sequence from init till end and push them in a channel
func getNum(init, end int, c chan int) {

	for i := init; i <= end; i++ {
		c <- i                                          // write value on the channel
		time.Sleep(time.Duration(2) * time.Microsecond) // introduce a timedelay to observe concurrency
	}
	close(c) // close the channel once the task is completed

}

// function to print Hello World repeatedly for num times
func printHello(num int) {
	for i := 0; i < num; i++ {
		fmt.Print("Hello World,")
		time.Sleep(time.Duration(10) * time.Microsecond) // introduce a timedelay to observe concurrency
	}
}
*/

package main

import (
	"fmt"
)

func senderTask(senderId int) {
	// Code to process a subtask
	fmt.Printf("Sender %v sent.\n", senderId)
}

func main() {

	numSenders := 10

	// Create a worker pool with a fixed number of goroutines
	chann := make(chan int, 5)

	/*
		// Start goroutines to process tasks in parallel
		for i := 0; i < numSenders; i++ {
			chann <- i // Add task to the worker pool
			go func(taskID int) {
				senderTask(taskID)
				for {
					str, ok := <-chann
					if !ok {
						break
					}
					fmt.Printf("Received %v\n", str)
				}
				<-chann // Remove task from the worker pool
			}(i)
		}
	*/

	// Start goroutines to process tasks in parallel
	for i := 0; i < numSenders; i++ {
		chann <- i // Add task to the worker pool
		go func(taskID int) {
			senderTask(taskID)

			<-chann // Remove task from the worker pool
		}(i)

		go func() {
			for {
				str, ok := <-chann
				if !ok {
					break
				}
				fmt.Printf("Received %v\n", str)
			}
		}()
	}

	// Wait for all numTask tasks to finish
	for i := 0; i < cap(chann); i++ {
		chann <- 0
	}
}

/* output:
Sender 4 sent.
Received 0
Received 1
Received 2
Received 3
Received 4
Received 5
Sender 9 sent.
Received 7
Received 8
Received 9
Received 0
Received 0
Received 0
*/
