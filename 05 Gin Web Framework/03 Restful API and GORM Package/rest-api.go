// using this api we will be aable to perform CRUD operations in the postgre database from a web-page

// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Create a Gin router instance
// 	router := gin.Default()

// 	// Serve the static files from the "static" directory
// 	router.Static("/crud", "./Rest-Api-CRUD")

// 	// Start the server
// 	router.Run(":8080")
// }

package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  int
}

func main() {

	dsn := "host=localhost user=postgres password=rabin@123 dbname=apple port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database", err)
	} else {
		fmt.Println("Successfully connected to database", db)
	}
	// Create

	// Migrate the schema, the below line of code basically creates an empty table

	// db.AutoMigrate(&User{})

	// Create a Gin router instance
	router := gin.Default()
	// router.Static("/", "./Rest-Api-CRUD")
	// Serve the form.html file when accessing the root route
	router.GET("/", func(c *gin.Context) {
		c.File("./Rest-Api-CRUD/index.html")
	})

	router.GET("/create", func(c *gin.Context) {
		c.File("./Rest-Api-CRUD/create.html")
	})

	router.GET("/insert", func(c *gin.Context) {
		c.File("./Rest-Api-CRUD/insert.html")
	})
	router.GET("/read", func(c *gin.Context) {
		c.File("./Rest-Api-CRUD/read.html")
	})
	router.GET("/update", func(c *gin.Context) {
		c.File("./Rest-Api-CRUD/update.html")
	})
	router.GET("/delete", func(c *gin.Context) {
		c.File("./Rest-Api-CRUD/delete.html")
	})

	// Define a route to handle POST requests
	router.POST("/insert-form", func(c *gin.Context) {
		// Read form data from the request
		name := c.PostForm("name")
		id := c.PostForm("id")
		age := c.PostForm("age")

		intAge, _ := strconv.Atoi(age)
		intId, _ := strconv.Atoi(id)

		user1 := User{ID: uint(intId), Name: name, Age: intAge}
		db.Create(&user1)

		// fmt.Println("Successfully inserted to database")
		fmt.Println("Name:", name)
		fmt.Println("Id:", id)
		fmt.Println("Age:", age)

	})

	router.POST("/update-form", func(c *gin.Context) {
		// Read form data from the request
		name := c.PostForm("name")
		id := c.PostForm("id")
		// age := c.PostForm("age")

		// intAge, _ := strconv.Atoi(age)
		intId, _ := strconv.Atoi(id)

		//update
		db.Model(User{}).Where("ID= ?", intId).Updates(User{Name: name})

		// fmt.Println("Successfully inserted to database")
		fmt.Println("Name:", name)
		fmt.Println("Id:", id)

	})

	router.POST("/read-form", func(c *gin.Context) {
		// Read form data from the request
		name := c.PostForm("name")
		id := c.PostForm("id")
		age := c.PostForm("age")

		intAge, _ := strconv.Atoi(age)
		intId, _ := strconv.Atoi(id)

		user1 := User{ID: uint(intId), Name: name, Age: intAge}
		db.Create(&user1)

		// fmt.Println("Successfully inserted to database")
		fmt.Println("Name:", name)
		fmt.Println("Id:", id)
		fmt.Println("Age:", age)

	})

	router.POST("/delete-form", func(c *gin.Context) {
		// Read form data from the request

		id := c.PostForm("id")
		intId, _ := strconv.Atoi(id)

		db.Where("ID = ?", intId).Delete(&User{}) // deletes data with Id = 4
		// fmt.Println("Successfully inserted to database")
		fmt.Println("Id:", id)

	})

	// Start the server
	router.Run(":3000")
}
