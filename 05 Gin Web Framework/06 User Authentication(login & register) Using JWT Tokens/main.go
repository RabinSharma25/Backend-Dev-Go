// here we will be setting jwt tokens

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type emp struct {
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func main() {

	dsn := "host=localhost user=postgres password=rabin@123 dbname=apple port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database", err)
	} else {
		fmt.Println("Successfully connected to database", db)
	}
	db.AutoMigrate(&emp{})
	router := gin.Default()

	// Set up routes
	router.POST("/register", registerHandler(db))
	router.POST("/login", loginHandler(db))

	router.Run(":8080")
}

func loginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user emp
		var userlogin emp // struct instance to store the temporary login data from user

		// Bind JSON data from the request body to the User struct
		// The code below will actually capture the data sent from postman
		if err := c.ShouldBindJSON(&userlogin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Log the received JSON data to the console
		fmt.Printf("Received JSON data from Postman: %+v\n", user)

		// token := "dummy-jwt-token"

		// Find the user by email
		if err := db.Where("email = ?", userlogin.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Compare the stored hash with the input password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userlogin.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Password Mismatch"})
			return
		}

		// Generate a JWT token with a 24-hour expiration time
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = user.Email
		claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

		// Sign the token with a secret key (change this to your actual secret key)
		secretKey := []byte("your-secret-key")
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Set the JWT token in a cookie that expires in 24 hours
		c.SetCookie("jwt", tokenString, int((24 * time.Hour).Seconds()), "/", "localhost", false, true)

		// Return the token as the response

		c.JSON(http.StatusOK, gin.H{"token": token})
		c.JSON(http.StatusOK, gin.H{"Login Successful": token})
		fmt.Println(token)

	}
}

func registerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user emp

		// Bind JSON data from the request body to the User struct
		// capture the data sent from postman
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Log the received JSON data to the console
		fmt.Printf("Received JSON data from Postman: %+v\n", user)

		// send a response
		c.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully",
		})

		// Hash the user's password before storing it in the database
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			// Handle the error
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"Error": "Failed to hash password"})
			return
		}

		// Create a new user record in the database
		newUser := emp{
			Email:    user.Email,
			Password: string(hashedPassword),
		}

		if err := db.Create(&newUser).Error; err != nil {
			// Handle the error (e.g., email already exists)
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"Error": "Email already registered"})
			return
		}

	}
}
