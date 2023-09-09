// here the cooking is set using the previous program

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Define a secret key for JWT (change this to your actual secret key)
var jwtSecret = []byte("your-secret-key") // this key must be saved in protected file in real applications

// AuthMiddleware is a middleware function to protect routes with JWT authentication

func main() {
	r := gin.Default()

	// Apply the AuthMiddleware to protect this route
	r.GET("/protected", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This route is protected"})
	})

	r.Run(":8080")
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the cookie
		tokenCookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check token validity
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired or invalid token"})
			c.Abort()
			return
		}

		// Token is valid; continue to the next middleware or controller
		c.Next()
	}
}

// Note: The jwt token will expire in 24 hrs...
// It can be set again by running the previous program i.e
//06 User Authentication(login & register) Using JWT Tokens/main.go
