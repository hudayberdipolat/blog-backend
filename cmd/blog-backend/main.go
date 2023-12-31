package main

import (
	categoryRoutes "github.com/hudayberdipolat/blog-backend/internal/categories/routes"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	userRoutes "github.com/hudayberdipolat/blog-backend/internal/user/routes"
	"github.com/hudayberdipolat/blog-backend/pkg/config"
	"github.com/hudayberdipolat/blog-backend/pkg/database"
	"github.com/hudayberdipolat/blog-backend/pkg/migrations"
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
		log.Println("Database connection error...")
		log.Fatal("error -->>", errDbConnection.Error())
	}
	log.Println("Database connection...")

	// migarate table begin

	if migrateError := migrations.MigrateTable(); migrateError != nil {
		log.Println("Table migration error...")
		log.Fatal("Error ---->>", migrateError.Error())
	}
	log.Println("Table migration...")
	// migrate table end

	// routes begin
	userRoutes.UserRoutes(app)
	categoryRoutes.CategoryRoutes(app)
	// routes end
	// run server
	port, ServerRunErr := config.ReadConfig()
	if ServerRunErr != nil {
		log.Println(ServerRunErr.Error())
	}
	log.Println("server running... ")
	log.Printf("RUN PORT:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Println("server run error")
		log.Fatal(err.Error())
		return
	}

}
