package main

import (
	"log"
	"github.com/vakshit/faq/internal/cmd"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	cmd.Execute()
}

// func main() {
// 	app := fiber.New()
// 	setupRoutes(app)
// 	// app.Get("/", helloWorld)
// 	database.ConnectMongo()

// 	app.Listen(":5000")
// 	fmt.Println("Server started on port 5000")

// }
