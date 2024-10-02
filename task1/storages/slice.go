package storages

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

func NewStorageSlice() *MyStorageSlice {
	return &MyStorageSlice{[]Book{}, map[int]int{}, 0}
}
