/*
// https://pkg.go.dev/time#Duration (func ParseDuration ¶)
package main

import (
	"fmt"
	"time"
)

func main() {
	second, _ := time.ParseDuration("1s")
	milis, _ := time.ParseDuration("1ms")
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	// The package also accepts the incorrect but common prefix u for micro.
	micro2, _ := time.ParseDuration("1us")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", hours.Seconds(), hours)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %dms in %v.\n", second.Milliseconds(), second)
	fmt.Printf("There are %dµs in %v.\n", milis.Microseconds(), milis)
	fmt.Printf("There are %dns in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %d nanoseconds in %v.\n", milis.Nanoseconds(), milis)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro)
}
*/
/*
Output:
10h0m0s
1h10m10s
There are 36000 seconds in 10h0m0s.
There are 4210 seconds in 1h10m10s.
There are 1000ms in 1s.
There are 1000µs in 1ms.
There are 1000ns in 1µs.
There are 1000000 nanoseconds in 1ms.
There are 1.00e-06 seconds in 1µs.
*/

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
*/
/*
Output:	(system local timestamp)
Current time:  2024-03-05 20:06:53.153144 +0700 +07 m=+1.001332000
Current time:  2024-03-05 20:06:54.153165 +0700 +07 m=+2.001338550
Current time:  2024-03-05 20:06:55.153172 +0700 +07 m=+3.001332343
Current time:  2024-03-05 20:06:56.153022 +0700 +07 m=+4.001168099
Current time:  2024-03-05 20:06:57.153217 +0700 +07 m=+5.001349046
Current time:  2024-03-05 20:06:58.153229 +0700 +07 m=+6.001347705
Current time:  2024-03-05 20:06:59.152205 +0700 +07 m=+7.000309761
Current time:  2024-03-05 20:07:00.153217 +0700 +07 m=+8.001308618
Current time:  2024-03-05 20:07:01.153276 +0700 +07 m=+9.001353330
Current time:  2024-03-05 20:07:02.15324 +0700 +07 m=+10.001303738
Done!
*/

// time's GoString() & Format(layout):
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	referenceDate := time.Date(2003, time.October, 9, 23, 0, 0, 0, time.UTC)
	fmt.Println("referenceDate:", referenceDate.Format("02/01/2006")) // Format the reference date as "dd/mm/yyyy"
	fmt.Println("referenceDate.Weekday():", referenceDate.Weekday())
	fmt.Println("now:", now.Format("02/01/2006")) // func (t Time) Format(layout string) string
	fmt.Println("now.Weekday():", now.Weekday())
	fmt.Println("now.Local():", now.Local()) // func (t Time) Local() Time

	yearsDiff := now.Year() - referenceDate.Year()
	monthsDiff := int(now.Month()) - int(referenceDate.Month())

	if monthsDiff < 0 {
		yearsDiff--
		monthsDiff += 12
	}

	fmt.Printf("%d years and %d months ago\n", yearsDiff, monthsDiff)

	// Compare compares the time instant t with u.
	// If t is before u, it returns -1; if t is after u,
	// it returns +1; if they're the same, it returns 0.
	fmt.Println("now.Compare(now+1ms):", now.Compare(now.Add(time.Millisecond*1))) // func (t Time) Compare(u Time) int
	fmt.Println("now.Compare(now):", now.Compare(now))

	end := time.Now()
	fmt.Println("Took:", end.Sub(now))
}
