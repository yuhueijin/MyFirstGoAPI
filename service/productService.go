package service

import (
	"github.com/yuhueijin/MyFirstGoAPI/repository"
)

func NewProductService() Service {
	return &service{product: repository.NewProduct()}
}

func (s *service) Add(name string) error {
	return s.product.Add(name)
}

func (s *service) Remove(id int) error {
	return s.product.Remove(id)
}

func (s *service) GetAll() ([]Model, error) {
	produts, err := s.product.GetAll()
	var results []Model
	for _, element := range produts {
		var newElement = Model{
			ID:   element.ID,
			Name: element.Name,
		}
		results = append(results, newElement)
	}
	return results, err
}
