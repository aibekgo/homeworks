package main

import (
	"fmt"
	"homeworks/lesson13/product"
)


func main() {
	p1 := product.Product{
		Name:  "p1",
		Price: 100,
	}
	p2 := product.Product{
		Name:  "p2",
		Price: 200,
	}
	p3 := product.Product{
		Name:  "p3",
		Price: 300,
	}
	basket1 := product.Basket{
		Name:     "b1",
		Products: []product.Product{p3, p2},
	}
	fmt.Println("---Before add---")
	basket1.ListProducts()
	basket1.AddProduct(p1)
	fmt.Println("---After add---")
	basket1.ListProducts()
}


/*
//Product:
Name:string
Price:int
GetFullInfo():string

//Basket:
Name:string
Products:[]Product

AddProduct(p Product)
PrintProducts() //вывод всех продуктов*/
