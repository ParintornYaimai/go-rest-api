package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ParintornYaimai/go-rest-api/api/handler"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	//สร้าง fiber
	app := fiber.New()

	connStr := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Cannot connet to DB:", err)
	}
	fmt.Println("Database connnectd!")

	defer db.Close()

	//health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("server is ok")
	})

	//route
	app.Get("/books", handler.GetAllBooks(db))

	log.Fatal(app.Listen(":3000"))
}
