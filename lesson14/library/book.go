package library

import "fmt"

type Book struct {
	Id		string
	Name 	string
	Author	string
	Price 	int
	Count	int
	Genre	Genre
}

func (b *Book) GetBookInfo() string {
	return fmt.Sprintf("Id:%s,Name:%s;Author:%s;Price:%d;Count:%d;Genre:%s", b.Id, b.Name,b.Author, b.Price, b.Count, b.Genre.Name)
}

func (b *Book) PrintBookInfo() {
	fmt.Println(b.GetBookInfo())
}

