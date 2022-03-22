package store

type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"`
}

type Store interface {
	Create(book *Book) error
	Update(id string, book *Book) error
	Get(id string) (Book, error)
	GetAll() ([]Book, error)
	Delete(id string) error
}
