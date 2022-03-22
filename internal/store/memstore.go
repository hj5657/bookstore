package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (m MemStore) Create(book *mystore.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Update(id string, book *mystore.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Get(id string) (mystore.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) GetAll() ([]mystore.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
