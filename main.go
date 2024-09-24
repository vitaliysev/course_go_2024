package main

import (
	library "awesomeProject/library"
	"fmt"
)

type Book = library.Book
type Library = library.Library
type Searcher = library.Searcher

func main() {
	books := []Book{
		{"Stephen King", "Different Seasons"},
		{"Dan Brown", "The Da Vinci Code"},
		{"J.K. Rowling", "Harry Potter and the Philosopher's Stone"},
	}
	var storage Searcher = library.NewStorageMap()
	var my_library Library = library.NewLibrary(library.IdSimple(), storage)
	my_library.Upload(books...)

	book, is_found := my_library.Search("Different Seasons")
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Println("Book not found")
	}

	book, is_found = my_library.Search("The Da Vinci Code")
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Println("Book not found")
	}

	my_library.ChangeIdGenerator(library.IdRandom())

	book, is_found = my_library.Search("Harry Potter and the Philosopher's Stone")
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Println("Book not found")
	}

	var storage2 Searcher = library.NewStorageSlice()
	my_library.SetStorage(storage2)

	books = []Book{
		{"Stephen King", "Different Seasons"},
		{"Dan Brown", "The Da Vinci Code"},
		{"J.K. Rowling", "Harry Potter and the Philosopher's Stone"},
	}

	my_library.Upload(books...)

	my_library.ChangeIdGenerator(library.IdSimple())
	book, is_found = my_library.Search("Different Seasons")
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Println("Book not found")
	}

	book, is_found = my_library.Search("The Da Vinci Code")
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Println("Book not found")
	}
}
