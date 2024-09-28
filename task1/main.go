package main

import (
	"fmt"
	"os"
	"task1/library"
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
	title := "Different Seasons"
	my_library.Upload(books...)

	book, is_found := my_library.Search(title)
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}
	title = "The Da Vinci Code"
	book, is_found = my_library.Search(title)
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}

	my_library.ChangeIdGenerator(library.IdRandom())

	title = "Harry Potter and the Philosopher's Stone"
	book, is_found = my_library.Search(title)
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
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
	title = "Different Seasons"
	book, is_found = my_library.Search(title)
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}
	title = "The Da Vinci Code"
	book, is_found = my_library.Search(title)
	if is_found {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}
}
