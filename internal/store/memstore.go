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

func (ms *MemStore) Create(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.books[book.Id]; ok {
		return mystore.ErrExist
	}

	nBook := *book
	ms.books[book.Id] = &nBook

	return nil
}

func (ms *MemStore) Update(id string, book *mystore.Book) error {
	//TODO implement me
	panic("implement me")
}

func (ms *MemStore) Get(id string) (mystore.Book, error) {
	ms.RLock()
	defer ms.RUnlock()

	t, ok := ms.books[id]
	if ok {
		return *t, nil
	}
	return mystore.Book{}, mystore.ErrNotFound
}

func (ms *MemStore) GetAll() ([]mystore.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (ms MemStore) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
