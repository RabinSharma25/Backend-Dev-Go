// CRUD opeations using native go driver for prosgreSQL i.e pq

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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

	res, err := db.Query("CREATE TABLE students(f_name varchar(30),l_name varchar(30),roll_no int , PRIMARY KEY(roll_no))")
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("Query executed successfully\n", res)
	}

	// Insert

	res, err = db.Query("INSERT INTO students VALUES('Ashwin','Adhikari',2)")

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("Query executed successfully", res)
	}

	// Read
	rows, err := db.Query("SELECT * FROM students")

	if err != nil {
		fmt.Println("Error getting data", err)
	}

	for rows.Next() {
		var fname string
		var lname string
		var id int
		if err := rows.Scan(&fname, &lname, &id); err != nil { // short hand notation
			fmt.Println("Error fetching data", err)
			continue
		}
		fmt.Println("The fetched data is", fname, lname, id)
		transact(db) // calling the transact function

	}

	// Update

	res, err = db.Query("UPDATE students set l_name = 'Nepal' 	WHERE roll_no = 1")

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("Query executed successfully\n", res)
	}

	// Delete

	res, err = db.Query("DELETE FROM students WHERE roll_no = 1")

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("Query executed successfully\n", res)
	}

}

// Transactions

// func transact(db *sql.DB) {

// 	tx, err := db.Begin()

// 	if err != nil {
// 		fmt.Println("Error tranacting", err)
// 		return
// 	}

// 	_, err = tx.Query("INSERT INTO students VALUES('Ashwin','Adhikari',4)")
// 	if err != nil {
// 		fmt.Println("Failed to insert", err)
// 		tx.Rollback()
// 		return
// 	}

// 	_, err = tx.Query("INSERT INTO students VALUES('Ashwin','Adhikari',7)")
// 	if err != nil {
// 		fmt.Println("Failed to insert", err)
// 		tx.Rollback()
// 		return
// 	}

// 	_, err = tx.Query("INSERT INTO students VALUES('Ashwin','Adhikari',9)")
// 	if err != nil {
// 		fmt.Println("Failed to insert", err)
// 		tx.Rollback()
// 		return
// 	}

// 	tx.Commit()

// }

// this is an alternative way to write the above function
func transact(db *sql.DB) {

	tx, err := db.Begin()

	if err != nil {
		fmt.Println("Error tranacting", err)
		return
	}
	defer tx.Rollback() // when ever execution is is completed(i.e a value is returned) this line gets executed
	_, err = tx.Query("INSERT INTO students VALUES('Bijay','Sharma',4)")
	if err != nil {
		fmt.Println("Failed to insert", err)
		// tx.Rollback()
		return
	}

	_, err = tx.Query("INSERT INTO students VALUES('Nisha','Chettri',7)")
	if err != nil {
		fmt.Println("Failed to insert", err)
		// tx.Rollback()
		return
	}

	_, err = tx.Query("INSERT INTO students VALUES('Aruna','Pradhan',9)")
	if err != nil {
		fmt.Println("Failed to insert", err)
		// tx.Rollback()
		return
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println("Failed to commit ", err)
	}

}
