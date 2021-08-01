package main

import (
	"fmt"
	"homeworks/lesson13/product"
)

func main() {
p1 := product.Product{
	Name:  "Maslo",
	Price: 650,
}
p2 := product.Product{
		Name:  "Hlebuwka",
		Price: 100,
}
b1 :=
fmt.Println(p1.GetFullInfo())
fmt.Println(p2.GetFullInfo())



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
