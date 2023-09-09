package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
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
	db.AutoMigrate(&User{})
	router := gin.Default()

	// Set up routes
	// TODO: Define routes for login, registration, and home pages
	router.POST("/register", registerHandler(db))
	router.POST("/login", loginHandler(db))
	router.GET("/home", homeHandler)
	router.GET("/register", getRegHandler)
	router.GET("/login", getLogHandler)
	router.Run(":8080")
}

func homeHandler(c *gin.Context) {
	c.File("home.html")
}

func getRegHandler(c *gin.Context) {
	c.File("register.html")
}

func getLogHandler(c *gin.Context) {
	c.File("login.html")
}

func loginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		var user User
		// Find the user by email
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Invalid credentials"})
			return
		}

		// Compare the stored hash with the input password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Invalid credentials"})
			return
		}

		// Successful login
		// TODO: Set a session or token to remember the user's login state

		// Redirect to the home page
		c.Redirect(http.StatusSeeOther, "/home")
	}
}

func registerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		// Hash the user's password before storing it in the database
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			// Handle the error
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"Error": "Failed to hash password"})
			return
		}

		// Create a new user record in the database
		newUser := User{
			Email:    email,
			Password: string(hashedPassword),
		}

		if err := db.Create(&newUser).Error; err != nil {
			// Handle the error (e.g., email already exists)
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"Error": "Email already registered"})
			return
		}

		// Successful registration
		// TODO: Set a session or token to remember the user's login state

		// Redirect to the home page or login page
		c.Redirect(http.StatusSeeOther, "/login")
	}
}
