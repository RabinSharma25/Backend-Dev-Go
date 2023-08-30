// CRUD opeations using native go driver for prosgreSQL

package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {

	// Connecting to database
	connStr := "user=postgres dbname=apple password=rabin@123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Error connecting to database", err)
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	// Note:- The db.Query() function returns two values the first value is the pointer to the data and the second value is the error message if any

	// Create
	{
		res, err := db.Query("CREATE TABLE students(f_name varchar(30),l_name varchar(30),roll_no int , PRIMARY KEY(roll_no))")
		if err != nil {
			fmt.Println("error", err)
		} else {
			fmt.Println("Query executed successfully\n", res)
		}
	}

	// Insert
	{
		res, err := db.Query("INSERT INTO students VALUES('Ashwin','Adhikari',2)")

		if err != nil {
			fmt.Println("error", err)
		} else {
			fmt.Println("Query executed successfully", res)
		}
	}

	// Read

	// Update
	{
		res, err := db.Query("UPDATE students set l_name = 'Nepal' 	WHERE roll_no = 1")

		if err != nil {
			fmt.Println("error", err)
		} else {
			fmt.Println("Query executed successfully\n", res)
		}
	}

	// Delete
	{
		res, err := db.Query("DELETE FROM students WHERE roll_no = 1")

		if err != nil {
			fmt.Println("error", err)
		} else {
			fmt.Println("Query executed successfully\n", res)
		}
	}

}
