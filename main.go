package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)
import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"time"
)

func main() {

	fmt.Println("Hello, World!")
	app := fiber.New()
	app.Get("/abc", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! 1234")
	})
	app.Post("/sum", AddOperation)
	app.Listen(":3000")

	if err := CreateDbObject(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

}

func AddOperation(c *fiber.Ctx) error {
	var request AddOperationRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := request.Num1 + request.Num2
	response := AddOperationResponse{Result: result}

	return c.JSON(response)
}

type AddOperationRequest struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type AddOperationResponse struct {
	Result int `json:"result"`
}

var DB *sql.DB

const (
	HOST     = "dpg-d1ibu62li9vc73fpqmmg-a"
	PORT     = "5432"
	USERNAME = "ajay_k045_user"
	PASSWORD = "JqfFofPysrAIp8aPst8OQLq8yQNQW5d3"
	DBNAME   = "Ajay"
)

func GetPsqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		HOST, PORT, USERNAME, PASSWORD, DBNAME)
}

func CreateDbObject() error {
	var err error

	DB, err = sql.Open("postgres", GetPsqlInfo())
	if err != nil {
		return fmt.Errorf("error opening DB: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %w", err)
	}

	fmt.Println("Connected successfully")

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(1 * time.Hour)

	return nil
}

// Continue with your logic here
