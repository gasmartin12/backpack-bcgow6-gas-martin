package users

import (
	"errors"
	"time"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class4-tt/internal/domain"
)

type Service interface {
	GetAll() ([]domain.User, error)
	Store(name string, lastName string, email string, age int, height int) (domain.User, error)
	Update(id int, name string, lastName string, email string, age int, height int, active bool) (domain.User, error)
	UpdateLastNameAndAge(id int, lastName string, age int) (domain.User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (s *service) Store(name string, lastName string, email string, age int, height int) (user domain.User, err error) {
	id, err := s.repository.LastId()
	if err != nil {
		err = errors.New("id could not be generated")
		return
	}
	id++
	creationDate := time.Now()
	user, err = s.repository.Store(id, name, lastName, email, age, height, true, creationDate)
	if err != nil {
		err = errors.New("user could not be created")
		return
	}
	return user, err
}

func (s *service) Update(id int, name string, lastName string, email string, age int, height int, active bool) (domain.User, error) {
	return s.repository.Update(id, name, lastName, email, age, height, active)
}
func (s *service) UpdateLastNameAndAge(id int, lastName string, age int) (domain.User, error) {
	return s.repository.UpdateLastNameAndAge(id, lastName, age)
}
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
