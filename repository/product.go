package repository

type model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product interface {
	Add(name string) error
	Remove(id int) error
	GetAll() ([]model, error)
}

type product struct{}