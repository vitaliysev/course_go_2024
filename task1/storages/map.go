package storages

import "task1/library"

type Book = library.Book
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

func NewStorageMap() *MyStorageMap {
	return &MyStorageMap{map[int]Book{}}
}
