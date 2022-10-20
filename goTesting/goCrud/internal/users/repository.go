package users

import (
	"fmt"
	"time"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goWeb/goCrud/pkg/store"
)

type User struct {
	Id           int
	Name         string
	LastName     string
	Email        string
	Age          int
	Height       int
	Active       bool
	CreationDate time.Time
}

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name string, lastName string, email string, age int, height int, active bool, creationDate time.Time) (User, error)
	LastId() (int, error)
	Update(id int, name string, lastName string, email string, age int, height int, active bool) (User, error)
	UpdateLastNameAndAge(id int, lastName string, age int) (User, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Read(&users)
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (r *repository) Store(id int, name string, lastName string, email string, age int, height int, active bool, creationDate time.Time) (User, error) {
	newUser := User{
		Id:           id,
		Name:         name,
		LastName:     lastName,
		Email:        email,
		Age:          age,
		Height:       height,
		Active:       active,
		CreationDate: creationDate,
	}
	var users []User
	if err := r.db.Read(&users); err != nil {
		return User{}, err
	}
	users = append(users, newUser)
	if err := r.db.Write(users); err != nil {
		return User{}, err
	}
	return newUser, nil
}

func (r *repository) LastId() (int, error) {
	var users []User
	if err := r.db.Read(&users); err != nil {
		return 0, err
	}
	if len(users) < 0 {
		return -1, nil
	}
	return users[len(users)-1].Id, nil
}

func (r *repository) Update(id int, name string, lastName string, email string, age int, height int, active bool) (User, error) {
	user := User{
		Id:       id,
		Name:     name,
		LastName: lastName,
		Email:    email,
		Age:      age,
		Height:   height,
		Active:   active,
	}
	var users []User
	if err := r.db.Read(&users); err != nil {
		return User{}, err
	}
	updated := false
	for i := range users {
		if users[i].Id == id {
			user.CreationDate = users[i].CreationDate
			users[i] = user
			updated = true
			break
		}
	}
	if !updated {
		return User{}, fmt.Errorf("user with id %d not found", id)
	}
	if err := r.db.Write(users); err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) UpdateLastNameAndAge(id int, lastName string, age int) (User, error) {
	updated := false
	var index int
	var users []User
	if err := r.db.Read(&users); err != nil {
		return User{}, err
	}
	for i := range users {
		if users[i].Id == id {
			users[i].LastName = lastName
			users[i].Age = age
			updated = true
			index = i
			break
		}
	}
	if !updated {
		return User{}, fmt.Errorf("user with id %d not found", id)
	}
	if err := r.db.Write(users); err != nil {
		return User{}, err
	}
	return users[index], nil
}

func (r *repository) Delete(id int) error {
	var users []User
	if err := r.db.Read(&users); err != nil {
		return err
	}
	found := false
	var index int
	for i := range users {
		if users[i].Id == id {
			index = i
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("user with id %d not found", id)
	}
	users = append(users[:index], users[index+1:]...)
	if err := r.db.Write(users); err != nil {
		return err
	}
	return nil
}
