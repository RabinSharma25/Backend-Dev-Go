package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router instance
	router := gin.Default()

	// Define a route to serve the static HTML file
	router.StaticFile("/static", "./static/index.html") // go to the /static route to acces the index.html page

	// Start the server
	router.Run(":8080")
}
