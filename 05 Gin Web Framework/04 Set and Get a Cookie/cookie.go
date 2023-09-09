package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router instance
	router := gin.Default()

	// Set a cookie
	router.GET("/setcookie", func(c *gin.Context) {
		// Set a cookie named "mycookie" with the value "Hello, Gin!" and other optional attributes
		c.SetCookie("mycookie", "Hello, Gin!", 10, "/", "localhost", false, true) // the numeric value shows the max-age of a cookie in seconds

		// Send a response
		c.String(http.StatusOK, "Cookie set successfully")
	})

	// Get a cookie
	router.GET("/getcookie", func(c *gin.Context) {
		// Get the value of the "mycookie" cookie
		myCookie, err := c.Cookie("mycookie")
		if err != nil {
			c.String(http.StatusNotFound, "Cookie not found")
		} else {
			c.String(http.StatusOK, "Value of mycookie: "+myCookie)
		}
	})

	// Start the Gin server
	router.Run(":8080")
}
