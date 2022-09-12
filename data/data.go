package data

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{ID: 1, Title: "The Hitchhiker's Guide to the Galaxy", Finished: false},
	{ID: 2, Title: "The Lord of the Rings", Finished: false},
	{ID: 3, Title: "The Hobbit", Finished: false},
	{ID: 4, Title: "The Catcher in the Rye", Finished: false},
	{ID: 5, Title: "The Grapes of Wrath", Finished: false},
	{ID: 6, Title: "To Kill a Mockingbird", Finished: false},
	{ID: 7, Title: "The Great Gatsby", Finished: false},
	{ID: 8, Title: "The Da Vinci Code", Finished: false},
	{ID: 9, Title: "Pride and Prejudice", Finished: false},
}

func findBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book

	m.RLock()
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.RUnlock()

	return index, book
}

func FinishedBook(id int, m *sync.RWMutex) {
	i, book := findBook(id, m)
	if i < 0 {
		return
	}
	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()

	fmt.Printf("Book %s finished\n", book.Title)
}
