package main

import (
	"fmt"
	"homeworks/lesson19/products"
	"log"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres" //postgres
	database := "go"
	password := "admin" //admin
	store, err := products.NewProductPostgresStore(host, database, user, password, port)
	if err != nil {
		log.Fatal(err)
		return
	}
	//u := &users.User{
	//	FullName: "Tleugazy Yerassyl",
	//	Username: "kirito",
	//	Password: "12345",
	//	Age:      23,
	//}
	//newUser, err := store.Create(u)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(newUser)
	products, err := store.List()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(products)

}
