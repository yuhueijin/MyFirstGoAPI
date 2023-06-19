package service

type Model struct {
   ID   int    `json:"id"`
   Name string `json:"name"`
}

type Service interface {
   Add(name string) error
   Remove(id int) error
   GetAll() ([]Model, error)
}

type svc struct{}

func NewService() Service {
   return &svc{}
}

func (s *svc) Add(name string) error {
	return nil
 }
 
 func (s *svc) Remove(id int) error {
	return nil
 }
 
 func (s *svc) GetAll() ([]Model, error) {
	return []Model{}, nil
 }