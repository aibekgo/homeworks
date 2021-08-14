package main

import (
	"fmt"
	"homeworks/lesson14/library"
)

func main() {
	genre1 := library.Genre{Name: "Fantastic"}
	genre2 := library.Genre{Name: "Horror"}
	genre3 := library.Genre{Name: "Love Story"}
	book1 := library.Book{
		Id:     "1",
		Name:   "Mars",
		Author: "Ilon Make",
		Price:  10,
		Count:  1,
		Genre:  genre1,
	}
	book2 := library.Book{
		Id:     "2",
		Name:   "Moon",
		Author: "Ilon Make",
		Price:  10,
		Count:  2,
		Genre:  genre1,
	}
	book3 := library.Book{
		Id:     "3",
		Name:   "Helen&Boris",
		Author: "Oleg Putin",
		Price:  20,
		Count:  3,
		Genre:  genre3,
	}
	book4 := library.Book{
		Id:     "4",
		Name:   "Covid",
		Author: "Flying Mouse",
		Price:  30,
		Count:  2,
		Genre:  genre2,
	}
	book5 := library.Book{
		Id:     "5",
		Name:   "Two parts",
		Author: "Olga Buzova",
		Price:  20,
		Count:  1,
		Genre:  genre3,
	}
	lib := library.Library{
		Name:     "Melowoman",
		Address:  "Gogol 55",
		Books: []library.Book{book1, book2, book3, book4},
	}
    lib.AddBook(book5)
	lib.PrintBooksInformation("add product result")
	lib.BuyBook("3", 1)
	lib.PrintBooksInformation("after buying of product by id 3")
	fmt.Println(lib.GetGenres())
	lib.PrintBooksCountByGenres()
	lib.PrintLibraryInfo()

}