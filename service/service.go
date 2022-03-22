package service

import (
	"bookstore/service/middleware"
	"bookstore/store"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	srv := &BookStoreServer{
		s: s,
		srv: &http.Server{
			Addr: addr,
		},
	}
	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.updateBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.getBookHandler).Methods("GET")
	router.HandleFunc("/book", srv.getAllBooksHandler).Methods("GET")
	router.HandleFunc("/book/{id}", srv.delBookHandler).Methods("DELETE")

	srv.srv.Handler = middleware.Logging(middleware.Validating(router))
	return srv
}

func (bs BookStoreServer) createBookHandler(writer http.ResponseWriter, request *http.Request) {
	dec := json.NewDecoder(request.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if err := bs.s.Create(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bs BookStoreServer) updateBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}
	dec := json.NewDecoder(request.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if err := bs.s.Update(id, &book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bs BookStoreServer) getBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}
	book, err := bs.s.Get(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	response(writer, book)
}

func response(writer http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bs BookStoreServer) getAllBooksHandler(writer http.ResponseWriter, _ *http.Request) {
	books, err := bs.s.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	response(writer, books)
}

func (bs BookStoreServer) delBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}
	err := bs.s.Delete(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bs BookStoreServer) ListenAndServe() (<-chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		err = bs.srv.ListenAndServe()
		errChan <- err
	}()
	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (bs BookStoreServer) Shutdown(ctx context.Context) error {
	err := bs.srv.Shutdown(ctx)
	return err
}
