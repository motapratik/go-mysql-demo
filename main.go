package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// selectQuery - Use to Call SQL Query
func selectQuery(db *sql.DB) {
	// Execute the query
	results, err := db.Query("SELECT id, name FROM country")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var c Country
		// for each row, scan the result into our tag composite object
		err = results.Scan(&c.ID, &c.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out ID and Name of Country
		fmt.Printf("ID=%d Name=%s\n", c.ID, c.Name)
	}
}

// selectProcedure - Use to Call Stored procedure
func selectProcedure(db *sql.DB) {
	// Execute the query
	results, err := db.Query("call GetCountryList()")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var c Country
		// for each row, scan the result into our tag composite object
		err = results.Scan(&c.ID, &c.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Printf("ID=%d Name=%s\n", c.ID, c.Name)
	}
}

// insertProcedure - Use to Call SQL Insert Stored Procedure
func insertProcedure(db *sql.DB, id int, name string) {
	// perform a db.Query insert
	insert, err := db.Query("call InsertCountry(?,?)", id, name)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func main() {

	// Open connection to database
	username := "User-Name"
	password := "Your-Password"
	hostname := "127.0.0.1" // for localhost
	port := "3306"
	dbname := "Database-Name"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	selectQuery(db)
	fmt.Printf("========\n")
	selectProcedure(db)
	fmt.Printf("========\n")
	insertProcedure(db, 4, "Germany") // for testing
	fmt.Printf("========\n")
	selectQuery(db)
}
