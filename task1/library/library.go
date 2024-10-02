package library

import (
	"math/rand"
	"time"
)

const sizeOfStorage = 256

type Book struct {
	Author string
	Title  string
}

type IdGenerator func() int

type Searcher interface {
	SearchInternal(int) (Book, bool)
	AddBook(int, Book)
	ClearInternal()
}

type Library interface {
	Search(string) (Book, bool)
	Upload(...Book)
	SetStorage(Searcher)
	ChangeIdGenerator(IdGenerator)
}

type MyLibrary struct {
	books   map[string]int
	makeID  IdGenerator
	storage Searcher
}

func (library *MyLibrary) SetStorage(newStorage Searcher) {
	library.storage = newStorage
	library.books = map[string]int{}
}

func (library *MyLibrary) Upload(books ...Book) {
	for _, book := range books {
		library.books[book.Title] = library.makeID()
		library.storage.AddBook(library.books[book.Title], book)
	}
}

func (library *MyLibrary) ChangeIdGenerator(newMakeID IdGenerator) {
	library.makeID = newMakeID
	booksTemp := []Book{}
	for Title, _ := range library.books {
		book, _ := library.Search(Title)
		booksTemp = append(booksTemp, book)
	}
	library.Clear()
	library.Upload(booksTemp...)
}
func (library *MyLibrary) Clear() {
	library.books = map[string]int{}
	library.storage.ClearInternal()
}

func (library MyLibrary) Search(Title string) (Book, bool) {
	id, ok := library.books[Title]
	if !ok {
		return Book{}, false
	}
	return library.storage.SearchInternal(id)
}

///

func IdSimple() func() int {
	nextID := 0
	return func() int {
		nextID++
		return nextID
	}
}

func IdRandom() func() int {
	rand.Seed(time.Now().UnixNano())
	used := [sizeOfStorage]bool{}
	return func() int {
		var i = 0
		for ; used[i]; i = rand.Intn(sizeOfStorage) {
		}
		used[i] = true
		return i
	}
}

func NewLibrary(makeID IdGenerator, storage Searcher) *MyLibrary {
	return &MyLibrary{map[string]int{}, makeID, storage}
}
