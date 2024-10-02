package main

import (
	"fmt"
	"os"
	"task1/library"
	"task1/storages"
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
	var storage Searcher = storages.NewStorageMap()
	var MyLibrary Library = library.NewLibrary(library.IdSimple(), storage)
	title := "Different Seasons"
	MyLibrary.Upload(books...)

	book, IsFound := MyLibrary.Search(title)
	if IsFound {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}
	title = "The Da Vinci Code"
	book, IsFound = MyLibrary.Search(title)
	if IsFound {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}

	MyLibrary.ChangeIdGenerator(library.IdRandom())

	title = "Harry Potter and the Philosopher's Stone"
	book, IsFound = MyLibrary.Search(title)
	if IsFound {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}

	var storage2 Searcher = storages.NewStorageSlice()
	MyLibrary.SetStorage(storage2)

	books = []Book{
		{"Stephen King", "Different Seasons"},
		{"Dan Brown", "The Da Vinci Code"},
		{"J.K. Rowling", "Harry Potter and the Philosopher's Stone"},
	}

	MyLibrary.Upload(books...)

	MyLibrary.ChangeIdGenerator(library.IdSimple())
	title = "Different Seasons"
	book, IsFound = MyLibrary.Search(title)
	if IsFound {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}
	title = "The Da Vinci Code"
	book, IsFound = MyLibrary.Search(title)
	if IsFound {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	} else {
		fmt.Printf("Book %s not found", title)
		os.Exit(3)
	}
}
