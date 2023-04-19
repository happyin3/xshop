package main

import (
	"fmt"
	"log"

	"xshop/summer/product"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hello/:name", product.Hello)

	log.Fatal(app.Listen(":3000"))

	fmt.Println("xshop summer")
}
