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
	user1 := User{ID: 2, Name: "hello", Age: 24}
	user2 := User{ID: 3, Name: "Rabin", Age: 23}
	user3 := User{ID: 4, Name: "Kishan", Age: 38}

	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)

	// Read
	var data User
	db.First(&data, 2) // find product with integer primary key (id = 2)
	fmt.Println(data)

	// update
	result := db.Model(User{}).Where("ID= ?", 2).Updates(User{Name: "hello", Age: 18})

	fmt.Println("The effected rows are ", result.RowsAffected) // returns updated records count
	fmt.Println("Updatin error", result.Error)                 // returns updating error

	// delete
	db.Where("ID = ?", 4).Delete(&User{}) // deletes data with Id = 4
	db.Delete(&user3)                     // deletes data with Id = 4 same as above line

}
