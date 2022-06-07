package main

import (
	"fmt"
	"log"
	"os"
	"workspace/database"
	router "workspace/route"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	// Initialize db instance
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize app
	app := fiber.New()

	// Set route
	router.SetRoute(app, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
