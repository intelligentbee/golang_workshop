package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO test.hello(world) VALUES('hello world!')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * from test.hello")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			log.Fatal(err)
		}
		log.Printf("found row containing %q", s)
	}
	rows.Close()
}
