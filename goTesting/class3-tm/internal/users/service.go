package users

import (
	"errors"
	"time"
)

type Service interface {
	GetAll() ([]User, error)
	Store(name string, lastName string, email string, age int, height int) (User, error)
	Update(id int, name string, lastName string, email string, age int, height int, active bool) (User, error)
	UpdateLastNameAndAge(id int, lastName string, age int) (User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (s *service) Store(name string, lastName string, email string, age int, height int) (user User, err error) {
	id, err := s.repository.LastId()
	if err != nil {
		err = errors.New("no se pudo generar un Id")
		return
	}
	id++
	creationDate := time.Now()
	user, err = s.repository.Store(id, name, lastName, email, age, height, true, creationDate)
	if err != nil {
		err = errors.New("no se pudo crear el usuario")
		return
	}
	return user, err
}

func (s *service) Update(id int, name string, lastName string, email string, age int, height int, active bool) (User, error) {
	return s.repository.Update(id, name, lastName, email, age, height, active)
}
func (s *service) UpdateLastNameAndAge(id int, lastName string, age int) (User, error) {
	return s.repository.UpdateLastNameAndAge(id, lastName, age)
}
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}