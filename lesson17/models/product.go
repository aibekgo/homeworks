package models

import (
	"database/sql"
	"log"
)
/*обявляем структуру*/
type Product struct {
	Id   string
	Name string
	Price int
}

/*функция insert в базу*/
func InsertData( nameValue string, priceValue int, db *sql.DB) {
	queryInsert := "insert into product ( name, price) values ( $1, $2)"
	result, err := db.Exec(queryInsert, nameValue, priceValue)
	if err != nil {
		log.Fatal(err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	if n <= 0 {
		log.Fatal(err)
		return
	}
}
/*функция select из базы*/
func GetData(db *sql.DB) []Product {
	products := []Product{}
	selectQuery := "select id, name, price from product"
	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for rows.Next() {
		item := Product{}
		err = rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		products = append(products, item)
	}
	return products
}
