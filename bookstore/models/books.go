package models

import (
	"database/sql"
	"errors"
	"golang.org/x/net/context"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks(ctx context.Context) ([]*Book, error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return nil, errors.New("models: could not get database connection pool from context")
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}