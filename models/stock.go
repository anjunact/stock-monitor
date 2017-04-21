package models

import (
	"time"
	"database/sql"

	_ "github.com/lib/pq"

)

type  Stock struct {
	Id  int
	Name string
	Code string
	Price float64
	Updated time.Time
}
var Db *sql.DB
func  init()  {
	var err error
	Db, err = sql.Open("postgres", "user=stock dbname=stock password=stock sslmode=disable")
	if err != nil {
		panic(err)
	}
}
func Stocks(limit int) (stocks []Stock, err error) {
	rows, err := Db.Query("select id, content, author from stocks limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		stock := Stock{}
		err = rows.Scan(&stock.Id, &stock.Name, &stock.Code,&stock.Price,&stock.Updated)
		if err != nil {
			return
		}
		stocks = append(stocks, stock)
	}
	rows.Close()
	return
}

func GetStock(code string) (stock Stock, err error) {
	stock = Stock{}
	err = Db.QueryRow("select id, name, code,price,updated from stocks where code = $1", code).Scan(&stock.Id, &stock.Name, &stock.Code,&stock.Price,&stock.Updated)
	return
}
func (stock *Stock) Create() (err error) {
	statement := "insert into stocks ( name,code,price) values ($1, $2,$3) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow( stock.Name,stock.Code,stock.Price).Scan(&stock.Id)
	return
}

func (stock *Stock) Update() (err error) {
	_, err = Db.Exec("update stocks set price = $2,name=$3 where id = $1", stock.Id, stock.Price,stock.Name)
	return
}

func (stock *Stock) Delete() (err error) {
	_, err = Db.Exec("delete from stocks where id = $1", stock.Id)
	return
}

func DeleteAll() (err error) {
	_, err = Db.Exec("delete from stocks")
	return
}
