// Concurrent Example (Single-threaded):
/*
package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d started\n", id)
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go task(i)
	}

	// Wait for a while to allow tasks to complete
	time.Sleep(3 * time.Second)
}
*/

// Concurrent Example (Single-threaded):
/*
package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d started\n", id)
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go task(i)
	}

	// Wait for a while to allow tasks to complete
	time.Sleep(3 * time.Second)
}
*/

/*
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started\n", id)
	rand.Seed(time.Now().UnixNano())                                   // Seed the random number generator
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond // Generate a random sleep duration between 0 and 1000 milliseconds
	time.Sleep(sleepDuration)                                          // Sleep for the random duration
	fmt.Printf("Task %d completed after %v\n", id, sleepDuration)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // Wait for all tasks to complete
}
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task(id int, startCh chan struct{}, wg *sync.WaitGroup) {
	<-startCh // Wait for signal to start
	defer wg.Done()
	fmt.Printf("Task %d started\n", id)

	rand.Seed(time.Now().UnixNano())
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepDuration)

	fmt.Printf("Task %d completed after %v\n", id, sleepDuration)
}

func main() {
	var wg sync.WaitGroup
	startCh := make(chan struct{}, 1) // Buffered channel to coordinate start of tasks

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(i, startCh, &wg)
	}

	// Signal all tasks to start concurrently
	close(startCh)

	wg.Wait() // Wait for all tasks to complete
}
