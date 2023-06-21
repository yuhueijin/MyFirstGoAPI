package service

import (
	"github.com/yuhueijin/MyFirstGoAPI/repository"
)

type Model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Service interface {
	Add(name string) error
	Remove(id int) error
	GetAll() ([]Model, error)
}

type service struct{ product repository.Product }
