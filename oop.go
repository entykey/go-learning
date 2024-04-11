/*
package main

import "fmt"

// Base class

	type Animal struct {
		Name string
	}

// Method for Animal

	func (a *Animal) Speak() {
		fmt.Println("Animal speaks")
	}

// Method for Animal that can be optionally overridden

	func (a *Animal) CustomBehavior() {
		fmt.Println("Default behavior for Animal")
	}

// Subclass

	type Dog struct {
		Animal // Embedding Animal in Dog
		Breed  string
	}

// Method for Dog

	func (d *Dog) Speak() {
		fmt.Println("Woof!")
	}

// Override CustomBehavior for Dog
// func (d *Dog) CustomBehavior() {
// 	fmt.Println("Custom behavior for Dog")
// }

// Subclass

	type Cat struct {
		Animal // Embedding Animal in Cat
		Breed  string
	}

// Method for Cat

	func (c *Cat) Speak() {
		fmt.Println("Meow!")
	}

// Override CustomBehavior for Cat

	func (c *Cat) CustomBehavior() {
		fmt.Println("Custom behavior for Cat")
	}

	func main() {
		// Create an instance of Dog
		dog := Dog{
			Animal: Animal{Name: "Buddy"},
			Breed:  "Golden Retriever",
		}

		// Call Speak method of Dog
		dog.Speak() // This will call the Speak method of Dog

		// Create an instance of Cat
		cat := Cat{
			Animal: Animal{Name: "Whiskers"},
			Breed:  "Siamese",
		}

		// Call Speak method of Cat
		cat.Speak() // This will call the Speak method of Cat

		// Call Speak method of Animal (embedded)
		// dog.Animal.Speak() // This will call the Speak method of Animal
		// cat.Animal.Speak() // This will call the Speak method of Animal

		// Call CustomBehavior method of Dog
		dog.CustomBehavior() // This will call the overridden CustomBehavior method for Dog

		// Call CustomBehavior method of Cat
		cat.CustomBehavior() // This will call the overridden CustomBehavior method
	}
*/

// Memory Access Synchronization
package main

import "fmt"

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

/*
package main

import "fmt"

func main() {
	var data int
	go func() { data++ }()
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}
*/

// data race free
/*
package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mutex = sync.Mutex{}

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
}

func main() {
	// Number of goroutines to run
	numGoroutines := 1000

	// Wait group to synchronize the termination of all goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Spawn multiple goroutines to increment the counter concurrently
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
*/

/*
package main

import (
	"fmt"
	"sync"
)

var counter = 0

func increment(wg *sync.WaitGroup) {
	counter++
	wg.Done()
}

func main() {
	// Number of goroutines to run
	numGoroutines := 1000

	// Wait group to synchronize the termination of all goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Spawn multiple goroutines to increment the counter concurrently
	for i := 0; i < numGoroutines; i++ {
		go increment(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}*/
