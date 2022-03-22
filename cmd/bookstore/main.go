package main

import (
	_ "bookstore/internal/store"
	"bookstore/store/factory"
)

func main() {
	_, err := factory.New("mem")
	if err != nil {
		panic(err)
	}
}
