package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router instance
	router := gin.Default()

	// Serve the static files from the "static" directory
	router.Static("/static", "./static")

	// Start the server
	router.Run(":8080")
}
