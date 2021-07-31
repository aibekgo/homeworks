package main

import "fmt"

type Book struct {
	Name     string
	Price    float64
	Category Category
}
type Category struct {
	Name string
}

func (c *Category) GetCategory() string {
s := fmt.Sprintf("%s", c.Name)
return s
}

func (b *Book) PrintData()  {
	fmt.Printf("BookName: %s BookPrice:%f BookCategory:%s",
	b.Name, b.Price, b.Category.GetCategory())

}

func main() {
	b1 := Book{
		Name:  "Les",
		Price:  10,
		Category: Category{Name: "Horor"}}

	b1.PrintData()
	}

	/*
	   Есть книга и есть категориии
	   Категории:
	           name:string

	   Book:
	       name:string
	       price:float64
	       category:Категории

	   printData()

	   BookName: , BookPrice: , BookCategory: ,*/
