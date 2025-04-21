package main

import (
	"share-kan/config"
	"share-kan/handler"

	"github.com/gofiber/fiber/v2"
)




func main() {
	//setup S3 Neva Object
	config.SetupNevaObject()
	
	//init Fiber
	app := fiber.New()

	//Endpoint File Management
	app.Post("/upload", handler.UploadHandler)


	//Run App
	app.Listen(":8080")
}