package main

import (
	"fmt"

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
	db.AutoMigrate(&User{})

	// Insert
	user1 := User{ID: 2, Name: "Ashwin", Age: 24}
	user2 := User{ID: 3, Name: "Rabin", Age: 23}
	user3 := User{ID: 4, Name: "Kishan", Age: 38}

	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)

	// Read
	var data User
	db.First(&data, 1) // find product with integer primary key
	fmt.Println(user1)

	// update

}
