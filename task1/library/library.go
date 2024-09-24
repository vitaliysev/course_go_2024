package library

import (
	"math/rand"
	"time"
)

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

//
type MyStorageSlice struct {
	internal []Book
	helper   map[int]int
	size     int
}

func (storage *MyStorageSlice) AddBook(id int, book Book) {
	storage.helper[id] = storage.size
	storage.size++
	storage.internal = append(storage.internal, book)
}

func (storage MyStorageSlice) SearchInternal(id int) (Book, bool) {
	book := storage.internal[storage.helper[id]]
	return book, true
}

func (storage *MyStorageSlice) ClearInternal() {
	storage.helper = map[int]int{}
	storage.size = 0
	storage.internal = []Book{}
}

///

///
type MyStorageMap struct {
	Internal map[int]Book
}

func (storage *MyStorageMap) AddBook(id int, book Book) {
	storage.Internal[id] = book
}

func (storage MyStorageMap) SearchInternal(id int) (Book, bool) {
	book, ok := storage.Internal[id]
	return book, ok
}

func (storage *MyStorageMap) ClearInternal() {
	storage.Internal = map[int]Book{}
}

///

///
type MyLibrary struct {
	books   map[string]int
	MakeId  IdGenerator
	storage Searcher
}

func (library *MyLibrary) SetStorage(newStorage Searcher) {
	library.storage = newStorage
	library.books = map[string]int{}
}

func (library *MyLibrary) Upload(books ...Book) {
	for _, book := range books {
		library.books[book.Title] = library.MakeId()
		library.storage.AddBook(library.books[book.Title], book)
	}
}

func (library *MyLibrary) ChangeIdGenerator(NewMakeId IdGenerator) {
	library.MakeId = NewMakeId
	books_temp := []Book{}
	for Title, _ := range library.books {
		book, _ := library.Search(Title)
		books_temp = append(books_temp, book)
	}
	library.Clear()
	library.Upload(books_temp...)
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
	next_id := 0
	return func() int {
		next_id++
		return next_id
	}
}

func IdRandom() func() int {
	rand.Seed(time.Now().UnixNano())
	used := [256]bool{}
	return func() int {
		var i = 0
		for ; used[i]; i = rand.Intn(256) {
		}
		used[i] = true
		return i
	}
}

func NewLibrary(MakeId IdGenerator, storage Searcher) *MyLibrary {
	return &MyLibrary{map[string]int{}, MakeId, storage}
}
func NewStorageMap() *MyStorageMap {
	return &MyStorageMap{map[int]Book{}}
}
func NewStorageSlice() *MyStorageSlice {
	return &MyStorageSlice{[]Book{}, map[int]int{}, 0}
}
