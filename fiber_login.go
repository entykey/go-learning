// LoginAsync -> ( getAllRoleNamesOfAUserByUserId(), getDepartmentsOfAUser() ) example
// go's sync.WaitGroup take care of the asyn/await so we no longer need them in go routine

// Go-LinQ example
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/ahmetb/go-linq"
)

// Department represents a department.
type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserData represents the user data.
type UserData struct {
	Id          string       `json:"id"`
	Roles       []string     `json:"roles"`
	Departments []Department `json:"departments"`
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func getAllRoleNamesOfAUserByUserId(userID string) []string {
	// Simulate delay
	time.Sleep(60 * time.Millisecond) // Millisecond

	// Fake result
	return []string{"Admin", "Member", "Moderator"}
}

func getDepartmentsOfAUser(userID string) []Department {
	// Simulate delay
	time.Sleep(120 * time.Millisecond)

	// Fake result
	return []Department{
		{ID: 1, Name: "IT"},
		{ID: 2, Name: "Sales"},
		{ID: 3, Name: "Marketing"},
	}
}

func main() {
	// timer
	start := time.Now()

	userID := "123"

	// waitgroup
	var wg sync.WaitGroup
	wg.Add(2)	// number of go routines

	var roles []string
	var departments []Department

	// Execute getAllRoleNamesOfAUserByUserId concurrently
	go func() {
		defer wg.Done()
		roles = getAllRoleNamesOfAUserByUserId(userID)

		duration := time.Since(start)
		fmt.Printf("Roles obtained in %.4fms\n", float64(duration.Microseconds())/1000)
	}()

	// Execute getDepartmentsOfAUser concurrently
	go func() {
		defer wg.Done()
		departments = getDepartmentsOfAUser(userID)

		duration := time.Since(start)
		fmt.Printf("Departments obtained in %.4fms\n", float64(duration.Microseconds())/1000)
	}()

	wg.Wait()

	userData := UserData{
		Id:          userID,
		Roles:       roles,
		Departments: departments,
	}

	// Convert to JSON using json.Marshal func
	jsonData, err := json.Marshal(userData)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println("")
	fmt.Println(string(jsonData))

	products := []Product{
		{ID: 1, Name: "Product A", Price: 10.99},
		{ID: 2, Name: "Product B", Price: 20.99},
		{ID: 3, Name: "Product C", Price: 15.99},
		{ID: 4, Name: "Product D", Price: 12.99},
		{ID: 5, Name: "Product E", Price: 18.99},
	}

	fmt.Println("")
	// Example 1: Filtering (get the products with price > 15.0)
	result := linq.From(products).
		Where(func(p interface{}) bool {
			product := p.(Product)
			return product.Price > 15.0
		}).
		Select(func(p interface{}) interface{} {
			product := p.(Product)
			return product.Name
		}).
		Results()

	fmt.Println("Filtered products:")
	for _, item := range result {
		fmt.Println(item)
	}

	// Example 2: Aggregation (Sum)
	totalPrice := linq.From(products).
		Select(func(p interface{}) interface{} {
			product := p.(Product)
			return product.Price
		}).
		Aggregate(func(a, b interface{}) interface{} {
			return a.(float64) + b.(float64)
		})

	fmt.Printf("Total price: %.2f\n", totalPrice)
}