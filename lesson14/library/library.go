package library

import (
"fmt"
"strings"
)

type Library struct {
	Name string
	Address string
	Books []Book
}


func (l *Library) PrintLibraryInfo (header string) {
	fmt.Println(fmt.Sprintf("\t---%s---", strings.ToUpper(header)))
	fmt.Println("Name: ", l.Name, "Address: ", l.Address)
		}

func (l *Library) PrintBooksInformation(header string) {
	fmt.Println(fmt.Sprintf("\t---%s---", strings.ToUpper(header)))
	for i := 0; i < len(l.Books); i++ {
		l.Books[i].PrintBookInfo()
	}
}

func (l *Library) GetAveragePrice() int {
	n := len(l.Books)
	sumi := 0
	avg := 0
	for i := 0; i < n; i++ {
		sumi += l.Books[i].Price
	}
	avg = sumi / n
	return avg
}

func (l *Library) GetCountOfBooks() int {
	n := len(l.Books)
	return n
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(id string) {
	books := []Book{}
	for i := 0; i < l.GetCountOfBooks(); i++ {
		if l.Books[i].Id != id {
			books = append(books, l.Books[i])
		}
	}
	l.Books = books
}

func (l *Library) BuyBook(id string, count int) {
	for i := 0; i < l.GetCountOfBooks(); i++ {
		if l.Books[i].Id == id {
			l.Books[i].Count = l.Books[i].Count - count
		}
	}
}

func (l *Library) GetGenres() []string {
	genres := []string{}
	for i := 0; i < l.GetCountOfBooks(); i++ {
		isExist := false
		for j := 0; j < len(genres); j++ {
			if l.Books[i].Genre.Name == genres[j] {
				isExist = true
				break
			}
		}
		if !isExist {
			genres = append(genres, l.Books[i].Genre.Name)
		}
	}
	return genres
}

func (l *Library) PrintBooksCountByGenres() {
	genres := l.GetGenres()
	for j := 0; j < len(genres); j++ {
		count := 0
		for i := 0; i < l.GetCountOfBooks(); i++ {
			if l.Books[i].Genre.Name == genres[j] {
				count += 1
			}
		}
		fmt.Println(fmt.Sprintf("%s:%d",genres[j],count))
	}
}
