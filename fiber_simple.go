package main

import (
	"log"

	"github.com/gofiber/fiber/v3" // Installation is done using the go get command: go get github.com/gofiber/fiber/v3
)

// fiber v2 main
// func main() {
//     app := fiber.New()

//     app.Get("/", func (c *fiber.Ctx) error {
//         return c.SendString("Hello, World!")
//     })

//     log.Fatal(app.Listen(":3000"))
// }

// https://docs.gofiber.io/next/
// fiber v3 main (not released yet)
func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello from Fiber v3 !")
	})

	// with query parameter
	// GET http://localhost:3000/hello%20world
	app.Get("/:value", func(c fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	app.Listen(":3000")

	log.Fatal(app.Listen(":3000"))
}
