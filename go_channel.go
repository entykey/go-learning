/*
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个缓冲区大小为 1 的通道 (Create two channels with buffer size 1)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	// 向 ch1 发送数据的 goroutine (A goroutine that sends data to ch1)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	// 向 ch2 发送数据的 goroutine (A goroutine that sends data to ch2)
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()

	// 无限循环，等待从通道接收数据
	// Infinite loop, waiting to receive data from channel)
	for {
		select {
		case x := <-ch1:
			fmt.Println("Received from ch1:", x)
			return
		case x := <-ch2:
			fmt.Println("Received from ch2:", x)
			return
		default:
			// 如果两个通道都已满，则打印一条消息并等待 500 毫秒
			// If both channels are full, print a message and wait 500 milliseconds
			fmt.Println("All channels are full")
			time.Sleep(500 * time.Millisecond)
			break
		}
	}
}
// output:
// All channels are full
// All channels are full
// All channels are full
// Received from ch2: 2

// 作者: Pistachiout
// 链接: https://juejin.cn/post/7265939798794764322
// 来源：稀土掘金
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// 2.1.1 Use time.Sleep to make the coroutine sleep to ensure that concurrent sub-coroutines are completed
// The time package provides time-related functions, the most commonly used of which is the time.Sleep function, which can make the current Goroutine sleep for a period of time. By combining Goroutine and time.Sleep, we can achieve concurrent execution of coroutines.
/*
   package main
   import (
   	"fmt"
   	"time"
   )
   func main() {
   	go task("Task 1")  // 启动协程1 (Start coroutine 1)
   	go task("Task 2")  // 启动协程2 (Start coroutine 2)
   	// The main coroutine sleeps for a period of time to ensure that the coroutine has enough time to execute
   	time.Sleep(3 * time.Second)
   }
   func task(name string) {
   	for i := 0; i < 5; i++ {
   		// Print task name and current iteration value
   		fmt.Println(name+":", i)
   		time.Sleep(500 * time.Millisecond)
   	}
   }
*/

/*
   package main
   import (
   	"fmt"
   	"sync"
   )
   func main() {
   	var wg sync.WaitGroup
   	wg.Add(2) // Set the counter of the waiting group to 2, indicating that there are two coroutines that need to wait.
   	go func() {		// Start coroutine 1
   		defer wg.Done() // After the coroutine is completed, the Done method is called to reduce the counter of the waiting group.
   		task("Task 1")
   	}()
   	go func() {	// Start coroutine 2
   		defer wg.Done() // After the coroutine is completed, the Done method is called to reduce the counter of the waiting group.
   		task("Task 2")
   	}()
   	wg.Wait() // Wait for all coroutines to complete and then end the main coroutine.
   }
   func task(name string) {
   	for i := 0; i < 5; i++ {
   		fmt.Println(name+":", i) // Print task name and current iteration value
   	}
   }
*/

/*
   package main

   import (
   "fmt"
   "sync"

   // benchmark
   "time"
   )

   // producer 向通道发送数据 (producer sends data to the channel)
   func producer(ch chan<- int, id int) {
   	for i := 0; i < 5; i++ {
   	ch <- i * id
   	}
   }
   // consumer 从通道接收数据 (consumer receives data from channel)
   func consumer(ch <-chan int, id int) {
   	for i := range ch {
   	// fmt.Printf("消费者 %d 接收到数据: %d\n", id, i)
   	fmt.Printf("Consumer %d received data: %d\n", id, i)

   	}
   }

   func main() {
   	// timer
   	start := time.Now()

   	ch := make(chan int, 10)

   	// wg 用于等待所有协程完成
   	// (wg is used to wait for all coroutines to complete)
   	var wg sync.WaitGroup

   	// producerWg 用于等待所有生产者协程完成，根据该等待组判断何时关闭通道
   	// (producerWg is used to wait for all producer coroutines to complete, and determine when to close the channel based on the waiting group)
   	var producerWg sync.WaitGroup

   	// 启动多个生产者协程
   	// (Start multiple producer coroutines)
   	for i := 0; i < 3; i++ {
   		// Increment producer wait group counter
   		producerWg.Add(1)

   		// Increment total wait group counter
   		wg.Add(1)
   		go func() {
   		producer(ch, i+1)

   		// Decrement the total wait group counter
   		wg.Done()

   		// Decrement the producer wait group counter
   		producerWg.Done()
   	}()
   }

   	// 启动多个消费者协程 (Start multiple consumer coroutines)
   	for i := 0; i < 3; i++ {
   		// Increment total wait group counter
   		wg.Add(1)
   		go func() {
   			consumer(ch, i+1)

   			// Decrement the total wait group counter
   			wg.Done()
   		}()
   	}
   	// Wait for all producer coroutines to complete and close the channel
   	producerWg.Wait()
   	close(ch)
   	// Wait for all coroutines to complete
   	wg.Wait()

   	duration := time.Since(start)
   	fmt.Printf("⌛️ Execution time: %dµs (%.4fms) (%.4fs)\n", duration.Microseconds(), float64(duration.Microseconds())/1000, float64(duration.Seconds()))
   }

   // 作者：Pistachiout
   // 链接：https://juejin.cn/post/7265939798794764322
   // 来源：稀土掘金
   // 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/*
   // The go keyword can be used for anonymous functions, and it is very simple for goroutine to implement multiple concurrency. Start a goroutine as follows:
   package main

   import (
   "fmt"
   // "sync"

   // benchmark
   "time"
   )

   func main() {
   	for i := 0; i < 10; i++ {
   		go func(n int) {
   			fmt.Println("执行了:", n)
   		}(i)
   	}
   	fmt.Println("main done!")
   	// 这里睡一会，防止main结束后，func1 来不及运行
   	time.Sleep(time.Second)
   }

   // 作者：small_to_large
   // 链接：https://juejin.cn/post/7240248679515865147
   // 来源：稀土掘金
   // 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/*
   // chatgpt, Rust-like mpsc implementation in Go
   // We can achieve a similar multi-producer single-consumer pattern using channels. Channels are the primary means of communication and synchronization between goroutines in Go.
   package main

   import (
       "fmt"
       "sync"
   )

   func producer(name string, queue chan<- string, wg *sync.WaitGroup) {
       defer wg.Done() // This line ensures that wg.Done() is called when the function exits
       for i := 0; i < 5; i++ { // Loop 5 times
           queue <- fmt.Sprintf("%s: %d", name, i) // Send a formatted message to the queue channel
       }
   }

   func consumer(queue <-chan string, done chan<- bool) {
       for msg := range queue { // Loop until the queue channel is closed
           fmt.Println("Received:", msg) // Print the received message
       }
       done <- true // Send a signal to the 'done' channel to indicate that the consumer is finished
   }

   func main() {
       queue := make(chan string) // Create a buffered channel for communication between producers and consumer
       done := make(chan bool) // Create a channel for signaling when the consumer is done
       var wg sync.WaitGroup // Create a WaitGroup to synchronize the producers

       // Start the consumer goroutine
       go consumer(queue, done)

       // Start multiple producer goroutines
       producers := []string{"Producer 1", "Producer 2", "Producer 3"}
       for _, name := range producers { // Iterate over the list of producer names
           wg.Add(1) // Increment the WaitGroup counter
           go producer(name, queue, &wg) // Start a producer goroutine
       }

       // Wait for all producers to finish
       wg.Wait()

       // Close the queue channel to signal the consumer to stop
       close(queue)

       // Wait for the consumer to finish
       <-done
   }
*/

/*
   // Rust mpsc worker thread like implementation in Go:
   package main

   import (
   	"fmt"
   	"sync"
   	"time"
   )

   const (
   	NUM_WORKERS = 4
   	NUM_TASKS   = 20
   )

   func worker(id int, tasks <-chan int, results chan<- string, wg *sync.WaitGroup) {
   	defer wg.Done()

   	fmt.Printf("Worker %d is waiting for tasks.\n", id)
   	for task := range tasks {
   		fmt.Printf("Worker %d is processing task: %d\n", id, task)

   		// Simulate some work
   		time.Sleep(200 * time.Millisecond)

   		// Send the result back to the result channel
   		results <- fmt.Sprintf("Result of task %d: done", task)
   	}
   	fmt.Printf("Worker %d is exiting.\n", id)
   }

   func main() {
   	tasks := make(chan int)
   	results := make(chan string)
   	var wg sync.WaitGroup

   	// Create worker goroutines
   	for i := 0; i < NUM_WORKERS; i++ {
   		wg.Add(1)
   		go worker(i, tasks, results, &wg)
   	}

   	// Submit tasks to the worker pool
   	go func() {
   		for i := 0; i < NUM_TASKS; i++ {
   			tasks <- i
   		}
   		close(tasks)
   	}()

   	// Collect results from the workers
   	go func() {
   		wg.Wait()
   		close(results)
   	}()

   	// Print received results
   	for result := range results {
   		fmt.Println("Received result:", result)
   	}
   }
*/
/* Output:
   Worker 0 is waiting for tasks.
   Worker 0 is processing task: 0
   Worker 1 is waiting for tasks.
   Worker 3 is waiting for tasks.
   Worker 1 is processing task: 1
   Worker 2 is waiting for tasks.
   Worker 3 is processing task: 2
   Worker 2 is processing task: 3
   Worker 2 is processing task: 4
   Received result: Result of task 3: done
   Received result: Result of task 2: done
   Worker 3 is processing task: 5
   Worker 1 is processing task: 6
   Received result: Result of task 1: done
   Received result: Result of task 0: done
   Worker 0 is processing task: 7
   Worker 0 is processing task: 8
   Received result: Result of task 7: done
   Received result: Result of task 6: done
   Received result: Result of task 4: done
   Worker 2 is processing task: 9
   Worker 1 is processing task: 10
   Worker 3 is processing task: 11
   Received result: Result of task 5: done
   Worker 3 is processing task: 12
   Received result: Result of task 11: done
   Received result: Result of task 9: done
   Received result: Result of task 8: done
   Received result: Result of task 10: done
   Worker 0 is processing task: 14
   Worker 1 is processing task: 15
   Worker 2 is processing task: 13
   Worker 0 is processing task: 16
   Received result: Result of task 14: done
   Worker 2 is processing task: 17
   Received result: Result of task 13: done
   Received result: Result of task 15: done
   Worker 1 is processing task: 18
   Worker 3 is processing task: 19
   Received result: Result of task 12: done
   Worker 1 is exiting.
   Received result: Result of task 18: done
   Received result: Result of task 19: done
   Received result: Result of task 16: done
   Received result: Result of task 17: done
   Worker 2 is exiting.
   Worker 3 is exiting.
   Worker 0 is exiting.
*/

/*
// try to learn the channel & waitgroup myself
package main

import (
	"fmt"
	"time"
)

// single-threaded:
// func someTask(data int) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Task %d executed\n", data)
// }

// func someTask(id int, data chan int) {
// 	for taskId := range data {
// 		time.Sleep(1 * time.Second)
// 		fmt.Printf("Task %d executed, ", taskId, "data: ", data, "\n")
// 	}
// }

func someTask(id int, data chan int) {
	for taskId := range data {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker: %d executed Task %d\n", id, taskId)
	}
}

func main() {

	// create a channel:
	channel := make(chan int)
	// var wg sync.WaitGroup

	// create 10 workers to execute the task
	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go someTask(i, channel)

	}

	// Filling channel with 10 numbers to be executed
	for i := 0; i < 10; i++ {
		channel <- i
	}

	// time.Sleep(10 * time.Second)

	// Close the channel when all goroutines are finished
	// go func() {
	// 	wg.Wait()
	// 	close(channel)
	// }()
}
*/

/*
// https://stackoverflow.com/questions/53838998/proper-way-to-close-a-channel
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go func(c chan<- string, t int) {
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
			c <- strconv.Itoa(t) + " : Done " + strconv.Itoa(rand.Intn(3000))
		}(ch, i)
	}
	for val := range ch {
		fmt.Println(val)

	}
}
*/

// Output
// 7 : Done 2619
// 1 : Done 2055
// 9 : Done 1291
// 3 : Done 2361
// 5 : Done 4
// 8 : Done 377
// 0 : Done 824
// 4 : Done 2125
// 6 : Done 2553
// 2 : Done 2825
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /Users/user/Documents/go_testing/go_channel.go:459 +0x105
// exit status 2

// fix:
/*
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(c chan<- string, t int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			c <- strconv.Itoa(t) + " : Done " + strconv.Itoa(rand.Intn(100))
		}(ch, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)

	}
}
*/

/*
   package main

   import (
   	"fmt"
   	"sync"
   	"time"
   )

   // SafeCounter is safe to use concurrently.
   type SafeCounter struct {
   	mu sync.Mutex
   	v  map[string]int
   }

   // Inc increments the counter for the given key.
   func (c *SafeCounter) Inc(key string) {
   	c.mu.Lock()

   	// Lock so only one goroutine at a time can access the map c.v.
   	c.v[key]++
   	c.mu.Unlock()
   }

   // Value returns the current value of the counter for the given key.
   func (c *SafeCounter) Value(key string) int {
   	c.mu.Lock()

   	// Lock so only one goroutine at a time can access the map c.v.
   	defer c.mu.Unlock()
   	return c.v[key]
   }

   func main() {

   	c := SafeCounter{v: make(map[string]int)}
   	for i := 0; i < 1000; i++ {
   		go c.Inc("somekey")
   	}

   	time.Sleep(time.Second)
   	fmt.Println(c.Value("somekey"))
   }
*/

/*
   // https://yourbasic.org/golang/for-loop/
   // https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/
   // https://yourbasic.org/golang/string-functions-reference-cheat-sheet/

   package main

   import (
   	"fmt"
   	"string"
   )

   func main() {
   	var n1, n2, n3 int

   	fmt.Println("n1: ", n1, ", n2: ", n2, ", n3: ", n3)

   	n1 = 3
   	n2 = 7

   	n3 = n1 + n2

   	fmt.Println("n3: ", n3);

   	fmt.Println("n1: ", n1, ", n2: ", n2, ", n3: ", n3)

   	// Basic for-each loop (slice or array)
   	a := []string{"Foo", "Bar"}
   	for i, s := range a {
   		fmt.Println(i, s)
   	}

   	// String iteration: runes or bytes
   	for i, ch := range "日本語" {
   		fmt.Printf("%#U starts at byte position %d\n", ch, i)
   	}

   	const s = "日本語"
   	for i := 0; i < len(s); i++ {
   		fmt.Printf("%x ", s[i])
   	}

   	fmt.Println("")

   	// A Go range loop iterates over UTF-8 encoded characters (runes):
   	for i, ch := range "Japan 日本" {
   	fmt.Printf("%d:%q ", i, ch)
   }

   	fmt.Println("")

   	// Channel iteration
   	ch := make(chan int)
   	go func() {
   		ch <- 1
   		ch <- 2
   		ch <- 3
   		close(ch)
   	}()
   	for n := range ch {
   		fmt.Println(n)
   	}

   	var str1 = strings.Join([]string{"a", "b"}, ":")
   	fmt.Println(str1)

   }
*/

// https://stackoverflow.com/questions/52975260/concurrency-vs-parallelism-when-executing-a-goroutine

/*
package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		fmt.Println("here")
		time.Sleep(time.Millisecond * 100)
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}
*/

/*
Output:

here
0
here
1
here
1
here
2
here
3
here
5
here
8
here
13
here
21
here
34
*/

/*
// without sleep, it's unsynchronized goroutines & messed up in completely undefined order (real async parallel behavior)
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		fmt.Println("here")
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(10, c)
	for i := range c {
		fmt.Println(i)
	}
}
*/

/*
Output:

here
0
1
here
here
1
2
here
here
3
5
here
here
8
13
here
here
21
34
here
*/

/*
// https://stackoverflow.com/questions/25417961/order-of-execution-inside-go-routines?rq=3
package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1

    fmt.Println("Inside the fibonacci")

    for {
        select {
        case c <- x:
            fmt.Println("Inside the for, first case, before reassigning ",x,y)
            x, y = y, x+y
	    fmt.Println("Inside the for, first case, after reassigning ",x,y)
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    fmt.Println("Begin of Main")
    c := make(chan int)
    quit := make(chan int)
    fmt.Println("Before gonig to the func")
    go func() {
        fmt.Println("Inside go routine")
	fmt.Println("Inside go routine... again")
        for i := 0; i < 10; i++ {
            fmt.Println("Inside go routine and the for, before printing the channel")
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fmt.Println("Before calling to fibonacci")
    fibonacci(c, quit)
    fmt.Println("Closing")
}
*/

/*
package main

import (
    "fmt"
    "time"
)

func main() {
    go sheep(1)
    go sheep(2)
    time.Sleep(100000)
}

func sheep(i int) {
    for ; ; i += 2 {
        fmt.Println(i,"sheeps")
    }
}
*/

/*
package main

import (
	"fmt"
)
*/
// learn go routine: https://juejin.cn/post/6844903623667744781

// sync.WaitGroup synchronization mechanism:
/*
func main() {
	t0 := time.Now()

	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()

	t1 := time.Now()
	fmt.Printf("The operation took %v to run.\n", t1.Sub(t0))
}

//*/

// using channels and synchronization techniques:
// safe: 66.357µs, 46.386µs, 108.769µs, 45.646µs, 48.313µs (slightly faster)
/*
func sayHello(done chan struct{}) {
	fmt.Println("hello")
	close(done) // Close the channel to signal completion
}

func main() {
	t0 := time.Now()

	done := make(chan struct{})

	go sayHello(done)

	// Block until the done channel is closed
	<-done

	t1 := time.Now()
	fmt.Printf("The operation took %v to run.\n", t1.Sub(t0))
}
*/

/*
Output: hello
*/

/*
// use channels to communicate between routines:
func main() {
	intStream := make(chan int)
	close(intStream) //  without this -> deadlock
	integer, ok := <-intStream
	fmt.Printf("(%v): %v\n", ok, integer)
	// (false): 0
}
*/

/*
func main() {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	// for {
	// 	integer, ok := <-intStream
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("%v ", integer)
	// }
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}

	fmt.Printf("\n")

	// Output: 1 2 3 4 5
}
*/

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}
*/

/*
package main

import (
	"fmt"
	"time"
)

// Request represents a request to be processed.
type Request struct {
	ID int
}

// MaxOutstanding is the maximum number of concurrent requests allowed.
var MaxOutstanding = 10

// sem is a semaphore channel used for limiting concurrent requests.
var sem = make(chan int, MaxOutstanding)

// process simulates processing a request.
func process(r *Request) {
	fmt.Printf("Processing request %d...\n", r.ID)
	time.Sleep(1 * time.Second) // Simulating processing time
	fmt.Printf("Request %d processed.\n", r.ID)
}

// handle processes a request, limiting concurrent processing using a semaphore.
func handle(r *Request) {
	sem <- 1   // Wait for the activity queue to clear.
	process(r) // It may take a long time.
	<-sem      // Complete; enables the next request to run.
}

// Serve listens for requests on a channel and handles them concurrently.
func Serve(queue chan *Request) {
	for {
		req := <-queue
		go handle(req) // No need to wait for handle to finish.
	}
}

func main() {
	// Create a request queue channel.
	queue := make(chan *Request)

	// Start serving requests.
	go Serve(queue)

	// Enqueue some requests.
	for i := 1; i <= 20; i++ {
		queue <- &Request{ID: i}
	}

	// Wait for a while to allow processing.
	time.Sleep(15 * time.Second)
}

*/

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

// Request represents a request to be processed.
type Request struct {
	ID int
}

// MaxOutstanding is the maximum number of concurrent requests allowed.
var MaxOutstanding = 10

// sem is a semaphore channel used for limiting concurrent requests.
var sem = make(chan int, MaxOutstanding)

// wg is a WaitGroup used to signal when all requests have been processed.
var wg sync.WaitGroup

// process simulates processing a request.
func process(r *Request) {
	fmt.Printf("Processing request %d...\n", r.ID)
	time.Sleep(1 * time.Second) // Simulating processing time
	fmt.Printf("Request %d processed.\n", r.ID)
}

// handle processes a request, limiting concurrent processing using a semaphore.
func handle(r *Request) {
	defer wg.Done()
	sem <- 1   // Wait for the activity queue to clear.
	process(r) // It may take a long time.
	<-sem      // Complete; enables the next request to run.
}

// Serve listens for requests on a channel and handles them concurrently.
// func Serve(queue chan *Request) {
// 	for {
// 		req := <-queue
// 		wg.Add(1)
// 		go handle(req) // No need to wait for handle to finish.
// 	}
// }


// func Serve(queue chan *Request) {
// 	for req := range queue {
// 		sem <- 1
// 		go func() {
// 			process(req) // 这儿有 Bug，解释见下。
// 			<-sem
// 		}()
// 	}
// }

// func Serve(queue chan *Request) {
// 	for req := range queue {
// 		sem <- 1
// 		go func(req *Request) {
// 			process(req)
// 			<-sem
// 		}(req)
// 	}
// }

func Serve(queue chan *Request) {
	for req := range queue {
		req := req // Create a new instance of req for this Go process.
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

func main() {
	// Create a request queue channel.
	queue := make(chan *Request)

	// Start serving requests.
	go Serve(queue)

	// Enqueue some requests.
	for i := 1; i <= 20; i++ {
		queue <- &Request{ID: i}
	}

	// Wait for all requests to be processed.
	wg.Wait()

	fmt.Println("All requests processed. Exiting...")
}

*/

/*
package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

*/

/*
package main

import "fmt"

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func handle(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

func main() {
	queue := make(chan *Request)

	go handle(queue)

	request := &Request{args: []int{3, 4, 6}, f: sum, resultChan: make(chan int)}

	// Send request
	queue <- request

	// Waiting for response
	fmt.Printf("Sum: %d\n", <-request.resultChan)
}
*/

/*
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go func(c chan<- string, t int) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			c <- strconv.Itoa(t) + " : Done " + strconv.Itoa(rand.Intn(1000))
		}(ch, i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
*/

/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 9; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("hi number %d from the spawned thread!\n", i)
			fmt.Printf("bye number %d from the spawned thread!\n", i)
		}(i)
	}

	wg.Wait()
}
*/

/*
package main

import (
	"fmt"
	"time"
)

type process struct {
	param []int
}

func (p process) Do() interface{} {
	var result int
	for _, i := range p.param {
		result += i
	}
	return result
}

func main() {

	po := go_parallel.NewParallelObject()
	po.AppendProcess("", process{
		[]int{1, 2, 3, 4, 5},
	})
	po.AppendProcess("", process{
		[]int{5, 6, 7, 8, 9},
	})
	po.SetTimeout(5000 * time.Millisecond)

	data, isTimeout := po.Run()

	var result []int
	for _, d := range data {
		result = append(result, d.Data.(int))
	}

	fmt.Println("isTimeout = ", isTimeout)
	fmt.Println(result)
}

// 作者：小嘴叭叭儿
// 链接：https://juejin.cn/post/7227020874889691196
// 来源：稀土掘金
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	// 呼び出す関数の前にgoキーワードを付けて呼び出す
	go hello()
	time.Sleep(time.Second * 2)
}

func hello() {
	fmt.Println("Hello")
}
*/

/*
// deadlock, try to implement mpsc example
package main

import (
	"fmt"
	"sync"
)

func sender(id int, messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := fmt.Sprintf("hi from thread %d", id)
	messages <- msg
}

func receiver(messages <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range messages {
		fmt.Println("Got", msg)
	}
	fmt.Println("Receiver no longer receiving messages. Exiting program.")
}

func main() {
	messages := make(chan string)
	var wg sync.WaitGroup

	// Spawn sender goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sender(i, messages, &wg)
	}

	// Spawn receiver goroutine
	wg.Add(1)
	go receiver(messages, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks finished. Exiting program.")
}
*/

/*
// In this step, We can modify the logic to automatically terminate the program once all tasks (sender goroutines) have completed their work. We can achieve this by using a sync.WaitGroup to track the completion of all sender goroutines and then exit the program when they are all done.
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := fmt.Sprintf("hi from thread %d", id)
	messages <- msg
}

func receiver(messages <-chan string) {
	for msg := range messages {
		fmt.Println("Got", msg)
	}
	fmt.Println("Receiver no longer receiving messages. Exiting program.")
}

func main() {
	start := time.Now()

	messages := make(chan string)
	var wg sync.WaitGroup

	// Spawn sender goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sender(i, messages, &wg)
	}

	// Spawn receiver goroutine
	go receiver(messages)

	// Wait for all sender goroutines to finish
	wg.Wait()

	// Close the messages channel to signal the receiver to stop
	close(messages)

	// fmt.Println("All tasks finished. Exiting program.")
	duration := time.Since(start)
	fmt.Println("Go Execution time:", duration.Microseconds(), "µs")

}
*/

/*
// STATUS: deadlock
package main

import (
	"fmt"
	"sync"
)

func sender1(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("hi from task %d, goroutine sender1", i)
		messages <- msg
	}
}

func sender2(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 4; i <= 10; i++ {
		msg := fmt.Sprintf("hi from task %d, goroutine sender2", i)
		messages <- msg
	}
}

func receiver(messages <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range messages {
		fmt.Println("Got", msg)
	}
	fmt.Println("Receiver no longer receiving messages. Exiting program.")
}

func main() {
	messages := make(chan string)
	var wg sync.WaitGroup

	// Spawn sender1 goroutine
	wg.Add(1)
	go sender1(messages, &wg)

	// Spawn sender2 goroutine
	wg.Add(1)
	go sender2(messages, &wg)

	// Spawn receiver goroutine
	wg.Add(1)
	go receiver(messages, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	close(messages) // still deadlock

	fmt.Println("All tasks finished. Exiting program.")
}
*/

/*
// STATUS: worked but sometimes give panic: send on closed channel
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender1(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("hi from task %d, goroutine sender1", i)
		messages <- msg
	}
}

func sender2(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 4; i <= 10; i++ {
		msg := fmt.Sprintf("hi from task %d, goroutine sender2", i)
		messages <- msg
	}
	close(messages) // Close the channel after all messages are sent
}

func receiver(messages <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range messages {
		fmt.Println("Got", msg)
	}
	fmt.Println("Receiver no longer receiving messages. Exiting program.")
}

func main() {
	start := time.Now()

	messages := make(chan string)
	var wg sync.WaitGroup

	// Spawn receiver goroutine
	wg.Add(1)
	go receiver(messages, &wg)

	// Spawn sender1 goroutine
	wg.Add(1)
	go sender1(messages, &wg)

	// Spawn sender2 goroutine
	wg.Add(1)
	go sender2(messages, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks finished. Exiting program.")

	duration := time.Since(start)
	fmt.Println("Go Execution time:", duration.Microseconds(), "µs")
}

*/

/*
// still deadlock again
package main

import (
	"fmt"
	"sync"
)

func sender1(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("hi from task %d, goroutine sender1", i)
		messages <- msg
	}
}

func sender2(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 4; i <= 10; i++ {
		msg := fmt.Sprintf("hi from task %d, goroutine sender2", i)
		messages <- msg
	}
}

func receiver(messages <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range messages {
		fmt.Println("Got", msg)
	}
}

func main() {
	messages := make(chan string)
	var wg sync.WaitGroup

	// Wait group for sender goroutines
	var senderWG sync.WaitGroup

	// Spawn receiver goroutine
	wg.Add(1)
	go func() {
		receiver(messages, &wg)
		senderWG.Wait() // Wait for all sender goroutines to finish
		close(messages)
	}()

	// Spawn sender1 goroutine
	senderWG.Add(1)
	go sender1(messages, &senderWG)

	// Spawn sender2 goroutine
	senderWG.Add(1)
	go sender2(messages, &senderWG)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks finished. Exiting program.")
}
*/

/*
// STATUS: fixed, deadlock-freed
// I see, the issue is that the receiver goroutine exits
// after completing its work, causing the program to deadlock
// as the main goroutine waits for the receiver to finish.
// To fix this, we need to ensure that the receiver goroutine waits for all sender goroutines to finish before it exits.
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender1(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("hi from task %d, goroutine sender1", i)
		messages <- msg
	}
}

func sender2(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 4; i <= 10; i++ {
		msg := fmt.Sprintf("hi from task %d, goroutine sender2", i)
		messages <- msg
	}
}

func receiver(messages <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range messages {
		fmt.Println("Got", msg)
	}
}

func main() {
	start := time.Now()

	messages := make(chan string)
	var wg sync.WaitGroup

	// Wait group for sender goroutines
	var senderWG sync.WaitGroup

	// Spawn receiver goroutine
	wg.Add(1)
	go func() {
		receiver(messages, &wg)
	}()

	// Spawn sender1 goroutine
	senderWG.Add(1)
	go sender1(messages, &senderWG)

	// Spawn sender2 goroutine
	senderWG.Add(1)
	go sender2(messages, &senderWG)

	// Wait for all sender goroutines to finish
	senderWG.Wait()

	// Close the messages channel after all sender goroutines have completed their work
	close(messages)

	// Wait for the receiver goroutine to finish
	wg.Wait()

	fmt.Println("All tasks finished. Exiting program.")

	duration := time.Since(start)
	fmt.Println("Go Execution time:", duration.Microseconds(), "µs")
}

*/
/*
Output:
Got hi from task 4, goroutine sender2
Got hi from task 1, goroutine sender1
Got hi from task 5, goroutine sender2
Got hi from task 6, goroutine sender2
Got hi from task 7, goroutine sender2
Got hi from task 8, goroutine sender2
Got hi from task 9, goroutine sender2
Got hi from task 10, goroutine sender2
Got hi from task 2, goroutine sender1
Got hi from task 3, goroutine sender1
All tasks finished. Exiting program.
Go Execution time: 131 µs
*/

/*
// channel-eg1-deadlock.go
// status: deadlock
package main

import (
	"fmt"
	"sync"
	"time"
)

const numTasks = 3

func sender(senderID int, senderChannel chan string) {
	fmt.Printf("Sender %d is waiting for send...\n", senderID)
	message := "hi"
	fmt.Printf("Sender %d: Starting\n", senderID)

	// Simulate some work
	// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Send a message to indicate completion
	senderChannel <- message
	fmt.Printf("Sent message from sender %d\n", senderID)
	fmt.Printf("Sender %d exiting...\n", senderID)
}

func receiver(senderChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Receiver thread started.")

	// Receive messages and print them as they arrive
	for msg := range senderChannel {
		fmt.Printf("Received: %s\n", msg)
	}

	fmt.Println("Receiver: All messages received.")
}

func main() {
	// Calculate the time taken
	start := time.Now()

	// Create channel for communication between sender and receiver
	senderChannel := make(chan string, numTasks)

	// Create waitgroup to wait for receiver to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Start receiver goroutine
	go receiver(senderChannel, &wg)

	// Spawn sender goroutines
	for i := 0; i < numTasks; i++ {
		go sender(i, senderChannel)
	}

	// Wait for all sender goroutines to complete
	wg.Wait()

	// Close the channel to signal the receiver goroutine
	close(senderChannel)

	// Wait for the receiver goroutine to finish
	wg.Wait()

	// Stop timer
	elapsed := time.Since(start)

	fmt.Printf("took %s to execute\n", elapsed)
}
*/

/*
// fixed deadlock
package main

import (
	"fmt"
	"sync"
	"time"
)

const numTasks = 3

func sender(senderID int, senderChannel chan string, wgSender *sync.WaitGroup) {
	defer wgSender.Done()
	fmt.Printf("Sender %d is waiting for send...\n", senderID)
	message := fmt.Sprintf("This is sender %d", senderID) // Format the message with sender's ID
	fmt.Printf("Sender %d: Starting\n", senderID)

	// Simulate some work
	// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Send a message to indicate completion
	senderChannel <- message
	fmt.Printf("Sent message from sender %d\n", senderID)
	fmt.Printf("Sender %d exiting...\n", senderID)
}

func receiver(senderChannel chan string, wgReceiver *sync.WaitGroup) {
	defer wgReceiver.Done()
	fmt.Println("Receiver thread started.")

	// Receive messages and print them as they arrive
	for msg := range senderChannel {
		fmt.Printf("Received: %s\n", msg)
	}

	fmt.Println("Receiver: All messages received.")
}

func main() {
	// Calculate the time taken
	start := time.Now()

	// Create channel for communication between sender and receiver
	senderChannel := make(chan string, numTasks)

	// Create waitgroup for sender goroutines
	var wgSender sync.WaitGroup
	wgSender.Add(numTasks)

	// Create waitgroup for receiver goroutine
	var wgReceiver sync.WaitGroup
	wgReceiver.Add(1)

	// Start receiver goroutine
	go receiver(senderChannel, &wgReceiver)

	// Spawn sender goroutines
	for i := 0; i < numTasks; i++ {
		go sender(i, senderChannel, &wgSender)
	}

	// Wait for all sender goroutines to complete
	wgSender.Wait()

	// Close the channel to signal the receiver goroutine
	close(senderChannel)

	// Wait for the receiver goroutine to finish
	wgReceiver.Wait()

	// Stop timer
	elapsed := time.Since(start)

	fmt.Printf("took %s to execute\n", elapsed)
}

*/
/*
Output 1:
Receiver thread started.
Sender 0 is waiting for send...
Sender 0: Starting
Sent message from sender 0
Sender 0 exiting...
Received: This is sender 0
Sender 1 is waiting for send...
Sender 1: Starting
Sent message from sender 1
Sender 1 exiting...
Received: This is sender 1
Sender 2 is waiting for send...
Sender 2: Starting
Sent message from sender 2
Sender 2 exiting...
Received: This is sender 2
Receiver: All messages received.
took 204.691µs to execute

Output 2:
Receiver thread started.
Sender 2 is waiting for send...
Sender 2: Starting
Sent message from sender 2
Sender 2 exiting...
Received: This is sender 2
Sender 0 is waiting for send...
Sender 0: Starting
Sent message from sender 0
Sender 0 exiting...
Received: This is sender 0
Sender 1 is waiting for send...
Sender 1: Starting
Sent message from sender 1
Sender 1 exiting...
Received: This is sender 1
Receiver: All messages received.
took 297.152µs to execute

Output 3:
Receiver thread started.
Sender 1 is waiting for send...
Sender 1: Starting
Sent message from sender 1
Sender 1 exiting...
Received: This is sender 1
Sender 2 is waiting for send...
Sender 0 is waiting for send...
Sender 0: Starting
Sent message from sender 0
Sender 0 exiting...
Received: This is sender 0
Sender 2: Starting
Sent message from sender 2
Sender 2 exiting...
Received: This is sender 2
Receiver: All messages received.
took 207.463µs to execute

Output 4:
Sender 2 is waiting for send...
Sender 2: Starting
Sent message from sender 2
Sender 2 exiting...
Receiver thread started.
Received: This is sender 2
Sender 0 is waiting for send...
Sender 0: Starting
Sent message from sender 0
Sender 0 exiting...
Received: This is sender 0
Sender 1 is waiting for send...
Sender 1: Starting
Sent message from sender 1
Sender 1 exiting...
Received: This is sender 1
Receiver: All messages received.
took 159.881µs to execute
*/

/*
// channel-eg2.go
// now put into a loop (each sender sends N messages):
package main

import (
	"fmt"
	"sync"
	"time"
)

const numTasks = 3
const numMessagesPerSender = 3

func sender(senderID int, senderChannel chan string, wgSender *sync.WaitGroup) {
	defer wgSender.Done()
	fmt.Printf("Sender %d is waiting for send...\n", senderID)
	fmt.Printf("Sender %d: Starting\n", senderID)

	// loop:
	for i := 1; i <= numMessagesPerSender; i++ {
		message := fmt.Sprintf("This is sender %d, message %d", senderID, i) // Format the message with sender's ID

		// Simulate some work
		// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// Send a message to indicate completion
		senderChannel <- message
		fmt.Printf("Sent message from sender %d\n", senderID)
	}

	fmt.Printf("Sender %d exiting...\n", senderID)
}

func receiver(senderChannel chan string, wgReceiver *sync.WaitGroup) {
	defer wgReceiver.Done()
	fmt.Println("Receiver thread started.")

	// Receive messages and print them as they arrive
	for msg := range senderChannel {
		fmt.Printf("Received: %s\n", msg)
	}

	fmt.Println("Receiver: All messages received.")
}

func main() {
	// Calculate the time taken
	start := time.Now()

	// Create channel for communication between sender and receiver
	senderChannel := make(chan string, numTasks)

	// Create waitgroup for sender goroutines
	var wgSender sync.WaitGroup
	wgSender.Add(numTasks)

	// Create waitgroup for receiver goroutine
	var wgReceiver sync.WaitGroup
	wgReceiver.Add(1)

	// Start receiver goroutine
	go receiver(senderChannel, &wgReceiver)

	// Spawn sender goroutines
	for i := 0; i < numTasks; i++ {
		go sender(i, senderChannel, &wgSender)
	}

	// Wait for all sender goroutines to complete
	wgSender.Wait()

	// Close the channel to signal the receiver goroutine
	close(senderChannel)

	// Wait for the receiver goroutine to finish
	wgReceiver.Wait()

	// Stop timer
	elapsed := time.Since(start)

	fmt.Printf("took %s to execute\n", elapsed)
}

*/
/*
Output 1:
Receiver thread started.
Sender 2 is waiting for send...
Sender 2: Starting
Sent message from sender 2
Sent message from sender 2
Sent message from sender 2
Sender 2 exiting...
Sender 0 is waiting for send...
Sender 0: Starting
Sent message from sender 0
Sender 1 is waiting for send...
Sender 1: Starting
Received: This is sender 2, message 1
Received: This is sender 2, message 2
Received: This is sender 2, message 3
Received: This is sender 0, message 1
Received: This is sender 0, message 2
Received: This is sender 1, message 1
Sent message from sender 1
Sent message from sender 1
Sent message from sender 1
Sender 1 exiting...
Received: This is sender 1, message 2
Received: This is sender 1, message 3
Sent message from sender 0
Sent message from sender 0
Sender 0 exiting...
Received: This is sender 0, message 3
Receiver: All messages received.
took 239.907µs to execute

Output 2:
Sender 2 is waiting for send...
Sender 2: Starting
Sent message from sender 2
Sent message from sender 2
Sent message from sender 2
Sender 2 exiting...
Sender 0 is waiting for send...
Receiver thread started.
Received: This is sender 2, message 1
Received: This is sender 2, message 2
Received: This is sender 2, message 3
Sender 0: Starting
Sender 1 is waiting for send...
Sender 1: Starting
Sent message from sender 1
Sent message from sender 1
Sent message from sender 1
Sent message from sender 0
Sender 1 exiting...
Received: This is sender 1, message 1
Received: This is sender 1, message 2
Received: This is sender 1, message 3
Received: This is sender 0, message 1
Received: This is sender 0, message 2
Sent message from sender 0
Sent message from sender 0
Sender 0 exiting...
Received: This is sender 0, message 3
Receiver: All messages received.
took 187.563µs to execute

Output 3:
Receiver thread started.
Sender 2 is waiting for send...
Sender 2: Starting
Sender 0 is waiting for send...
Sent message from sender 2
Sent message from sender 2
Sent message from sender 2
Sender 2 exiting...
Sender 0: Starting
Sent message from sender 0
Sender 1 is waiting for send...
Sender 1: Starting
Received: This is sender 2, message 1
Received: This is sender 2, message 2
Received: This is sender 2, message 3
Received: This is sender 0, message 1
Received: This is sender 0, message 2
Received: This is sender 1, message 1
Sent message from sender 1
Sent message from sender 1
Sent message from sender 0
Sent message from sender 0
Sender 0 exiting...
Received: This is sender 1, message 2
Received: This is sender 1, message 3
Received: This is sender 0, message 3
Sent message from sender 1
Sender 1 exiting...
Receiver: All messages received.
took 265.203µs to execute
*/

// *
// Make a Go version for the Rust code (In Go, goroutines are used instead of threads. Goroutines are lightweight, managed by the Go runtime, and multiplexed onto a small number of OS threads. They are more efficient than traditional threads and are designed to handle concurrent tasks efficiently.)
// STATUS: deadlock since we forgot to signal stop at the end of tasks
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, senderChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Sender %d is waiting for send...\n", id)
	curThreadID := id
	for i := id*10 - 9; i <= id*10; i++ {
		msg := fmt.Sprintf("This is thread %v, message %d", curThreadID, i)
		senderChannel <- msg
		fmt.Printf("Sent message from sender %d\n", id)
	}
	fmt.Printf("Sender %d exiting...\n", id)
}

func receiver(senderChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Receiver thread started.")
	for msg := range senderChannel {
		fmt.Println("Got", msg)
	}
	fmt.Println("Receiver: All messages received.")
}

func main() {
	// Start timer
	start := time.Now()

	senderChannel := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go receiver(senderChannel, &wg)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sender(i, senderChannel, &wg)
	}

	wg.Wait()
	close(senderChannel)

	// Stop timer
	elapsed := time.Since(start)
	fmt.Printf("took %s to execute\n", elapsed)
}

//*/

/*
// STATUS: still deadlock , why ??
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, senderChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Sender %d is waiting for send...\n", id)
	curThreadID := id
	for i := id*10 - 9; i <= id*10; i++ {
		msg := fmt.Sprintf("This is thread %v, message %d", curThreadID, i)
		senderChannel <- msg
		fmt.Printf("Sent message from sender %d\n", id)
	}
	fmt.Printf("Sender %d exiting...\n", id)
}

func receiver(senderChannel chan string, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	fmt.Println("Receiver thread started.")
	for {
		select {
		case msg, ok := <-senderChannel:
			if ok {
				fmt.Println("Got", msg)
			} else {
				fmt.Println("Receiver: All messages received.")
				done <- true
				return
			}
		}
	}
}

func main() {
	// Start timer
	start := time.Now()

	senderChannel := make(chan string)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go receiver(senderChannel, &wg, done)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sender(i, senderChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(senderChannel)
	}()

	<-done // Wait for the receiver to finish

	// Stop timer
	elapsed := time.Since(start)
	fmt.Printf("took %s to execute\n", elapsed)
}
*/

/*
// STATUS: deadlock resolved, panic happened
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, senderChannel chan string, wg *sync.WaitGroup, closeOnce *sync.Once) {
	defer wg.Done()
	defer closeOnce.Do(func() { close(senderChannel) }) // Close the channel when all senders are done
	fmt.Printf("Sender %d is waiting for send...\n", id)
	curThreadID := id
	for i := id*10 - 9; i <= id*10; i++ {
		msg := fmt.Sprintf("This is thread %v, message %d", curThreadID, i)
		senderChannel <- msg
		fmt.Printf("Sent message from sender %d\n", id)
	}
	fmt.Printf("Sender %d exiting...\n", id)
}

func receiver(senderChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Receiver thread started.")
	for msg := range senderChannel {
		fmt.Println("Got", msg)
	}
	fmt.Println("Receiver: All messages received.")
}

func main() {
	// Start timer
	start := time.Now()

	senderChannel := make(chan string)
	var wg sync.WaitGroup
	var closeOnce sync.Once

	wg.Add(1)
	go receiver(senderChannel, &wg)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sender(i, senderChannel, &wg, &closeOnce)
	}

	wg.Wait()

	// Stop timer
	elapsed := time.Since(start)
	fmt.Printf("took %s to execute\n", elapsed)
}
*/

/*
// STATUS: damn still panic
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, senderChannel chan string, wg *sync.WaitGroup, closeOnce *sync.Once) {
	defer wg.Done()
	fmt.Printf("Sender %d is waiting for send...\n", id)
	curThreadID := id
	for i := id*10 - 9; i <= id*10; i++ {
		msg := fmt.Sprintf("This is thread %v, message %d", curThreadID, i)

		// Check if the channel is closed before sending data
		select {
		case <-closeOnce.Done():
			return
		default:
			senderChannel <- msg
			fmt.Printf("Sent message from sender %d\n", id)
		}
	}
	fmt.Printf("Sender %d exiting...\n", id)
}

func receiver(senderChannel chan string, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	fmt.Println("Receiver thread started.")
	receivedAll := false
	for !receivedAll {
		select {
		case msg, ok := <-senderChannel:
			if ok {
				fmt.Println("Got", msg)
			} else {
				receivedAll = true
				fmt.Println("Receiver: All messages received.")
				done <- true
			}
		}
	}
}

func main() {
	// Start timer
	start := time.Now()

	senderChannel := make(chan string)
	done := make(chan bool)
	var wg sync.WaitGroup
	var closeOnce sync.Once

	wg.Add(1)
	go receiver(senderChannel, &wg, done)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sender(i, senderChannel, &wg, &closeOnce)
	}

	wg.Wait()

	// Close the sender channel after all senders have finished
	closeOnce.Do(func() { close(senderChannel) })

	// Wait for the receiver to finish
	<-done

	// Stop timer
	elapsed := time.Since(start)
	fmt.Printf("took %s to execute\n", elapsed)
}
*/

/*
// resolved
package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(id int, senderChannel chan string, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	defer func() {
		fmt.Printf("Sender %d exiting...\n", id)
	}()
	fmt.Printf("Sender %d is waiting for send...\n", id)
	curThreadID := id
	for i := id*10 - 9; i <= id*10; i++ {
		msg := fmt.Sprintf("This is thread %v, message %d", curThreadID, i)
		senderChannel <- msg
		fmt.Printf("Sent message from sender %d\n", id)
	}
	done <- true
}

func receiver(senderChannel chan string, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	fmt.Println("Receiver thread started.")
	for {
		select {
		case msg, ok := <-senderChannel:
			if ok {
				fmt.Println("Got", msg)
			} else {
				fmt.Println("Receiver: All messages received.")
				done <- true
				return
			}
		}
	}
}

func main() {
	// Start timer
	start := time.Now()

	senderChannel := make(chan string)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go receiver(senderChannel, &wg, done)

	senderDone := make(chan bool)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sender(i, senderChannel, &wg, senderDone)
	}

	go func() {
		for i := 0; i < 3; i++ {
			<-senderDone
		}
		close(senderChannel)
	}()

	// Wait for receiver and all sender goroutines to finish
	<-done
	wg.Wait()

	// Stop timer
	elapsed := time.Since(start)
	fmt.Printf("took %s to execute\n", elapsed)
}
*/

// deadlock
/*
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numTasks    = 10
	numThreads  = 2
	messageSize = 50
)

var (
	wg              sync.WaitGroup
	senderChannel   = make(chan string, numTasks)
	receiverChannel = make(chan string, numTasks)
)

func sender(senderID int) {
	defer wg.Done()

	fmt.Printf("Sender %d is waiting for send...\n", senderID)

	message := "hi"
	fmt.Printf("Sender %d: Starting\n", senderID)

	// Simulate some work
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Send a message to indicate completion
	senderChannel <- message
	fmt.Printf("Sent message from sender %d\n", senderID)
	fmt.Printf("Sender %d exiting...\n", senderID)
}

func receiver() {
	defer wg.Done()

	fmt.Println("Receiver thread started.")

	// Receive messages and print them as they arrive
	receivedCount := 0
	for receivedCount < numTasks {
		msg := <-receiverChannel
		fmt.Printf("Received: %s\n", msg)
		receivedCount++
	}

	fmt.Println("Receiver: All messages received.")
}

func main() {
	wg.Add(numThreads + 1) // Add receiver thread to wait group

	// Start receiver thread
	go receiver()

	// Start sender threads
	for i := 1; i <= numThreads; i++ {
		go sender(i)
	}

	// Wait for all sender threads to finish
	wg.Wait()

	// Close channels
	close(senderChannel)
	close(receiverChannel)

	fmt.Println("All sender threads finished.")
}
*/

// fix
/*
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numTasks    = 10
	numThreads  = 2
	messageSize = 50
)

var (
	wg              sync.WaitGroup
	senderChannel   = make(chan string, numTasks)
	receiverChannel = make(chan string, numTasks)
)

func sender(senderID int) {
	defer wg.Done()

	fmt.Printf("Sender %d is waiting for send...\n", senderID)

	message := "hi"
	fmt.Printf("Sender %d: Starting\n", senderID)

	// Simulate some work
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Send a message to the receiver channel
	receiverChannel <- message
	fmt.Printf("Sent message from sender %d\n", senderID)
	fmt.Printf("Sender %d exiting...\n", senderID)
}

func receiver() {
	defer wg.Done()

	fmt.Println("Receiver thread started.")

	// Receive messages and print them as they arrive
	receivedCount := 0
	for receivedCount < numTasks {
		msg := <-senderChannel
		fmt.Printf("Received: %s\n", msg)
		receivedCount++
	}

	fmt.Println("Receiver: All messages received.")
}

func main() {
	wg.Add(numThreads + 1) // Add receiver thread to wait group

	// Start receiver thread
	go receiver()

	// Start sender threads
	for i := 1; i <= numTasks; i++ {
		go sender(i)
	}

	// Wait for all sender threads to finish
	wg.Wait()

	// Close channels
	close(senderChannel)
	close(receiverChannel)

	fmt.Println("All sender threads finished.")
}
*/
/*
Output:
Sender 3 is waiting for send...
Sender 3: Starting
Sender 10 is waiting for send...
Sender 10: Starting
Sender 4 is waiting for send...
Sender 4: Starting
Sender 5 is waiting for send...
Sender 5: Starting
Sender 6 is waiting for send...
Sender 6: Starting
Sender 7 is waiting for send...
Sender 7: Starting
Sender 8 is waiting for send...
Sender 8: Starting
Sender 9 is waiting for send...
Sender 9: Starting
Receiver thread started.
Sender 2 is waiting for send...
Sender 2: Starting
Sender 1 is waiting for send...
Sender 1: Starting
Sent message from sender 5
Sender 5 exiting...
Sent message from sender 3
Sender 3 exiting...
Sent message from sender 1
Sender 1 exiting...
All sender threads finished.
*/
