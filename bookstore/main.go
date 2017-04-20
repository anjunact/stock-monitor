package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
)

type Book struct {
	isbn  string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://stock-monitor:123456@192.168.56.101/stock-monitor?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}