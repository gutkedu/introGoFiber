package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gutkedu/introGoFiber/book"
	"github.com/gutkedu/introGoFiber/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBCoon, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database successfully opened")

	database.DBCoon.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated successfully")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBCoon.Close()
	setupRoutes(app)
	app.Listen(3000)
}
