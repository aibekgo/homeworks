package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"homeworks/lesson17/models"
	"log"
)

func main() {
	/*переменная запроса*/
	createTableQuery := `
		create table if not exists product(
			id text NOT NULL DEFAULT nextval('seq_prct') ,
			name text,
	     	price integer,
			primary key(id)
		);
	`
	/*коннект к базе*/
	host := "localhost"
	port := "5432"
	user := "postgres" //postgres
	database := "go"
	password := "admin" //admin
	params := "sslmode=disable"
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", user, password, host, port, database, params)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	/*выполнение запроса*/
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
		return
	}
	/*вызов функции insert в базу*/
	models.InsertData( "iphone", 100,  db)
	/*вызов функции select в базу*/
	products:= models.GetData(db)
	fmt.Println(products)
}

