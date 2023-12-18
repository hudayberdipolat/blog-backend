package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/pkg/config"
	"github.com/hudayberdipolat/blog-backend/pkg/database"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Blog",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		BodyLimit:    2 * 1024,
	})
	// database connection
	errDbConnection := database.DatabaseConnection()
	if errDbConnection != nil {
		log.Println("Database connection error!!!")
		log.Fatal("error -->>", errDbConnection.Error())
	}
	// run server
	port, ServerRunerr := config.ReadConfig()
	if ServerRunerr != nil {
		log.Println(ServerRunerr.Error())
	}
	if err := app.Listen(":" + port); err != nil {
		log.Println("server run error")
		log.Fatal(err.Error())
		return
	}
}
