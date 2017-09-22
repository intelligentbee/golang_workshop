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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, world varchar(50))")
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
		var id int
		var world string
		if err := rows.Scan(&id, &world); err != nil {
			log.Fatal(err)
		}
		log.Printf("found row containing %d, %q", id, world)
	}
	rows.Close()
}
