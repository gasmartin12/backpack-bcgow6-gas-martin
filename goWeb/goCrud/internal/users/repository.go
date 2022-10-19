package users

import (
	"fmt"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goWeb/goCrud/pkg/store"
)

type User struct {
	Id         int     `json:"Id"`
	Name       string  `form:"Name" json:"Name"`
	LastName   string  `form:"LastName" json:"LastName"`
	Email      string  `form:"Email" json:"Email"`
	Age        int     `form:"Age" json:"Age"`
	Height     float64 `form:"Height" json:"Height"`
	Active     bool    `form:"Active" json:"Active"`
	DateCreate string  `form:"DateCreate" json:"DateCreate"`
}

var users []User

type Repository interface {
	GetAllUsers() ([]User, error)
	GetOneUser()
	SaveUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error)
	LastId() (int, error)
	UpdateUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error)
	DeleteUser(Id int) error
	UpdatePatch(Id int, LastName string, Age int) (User, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error) {

	var us []User
	err := r.db.Read(&us)
	if err != nil {
		return User{}, err
	}
	u := User{Id, Name, LastName, Email, Age, Height, Active, DateCreate}
	us = append(us, u)
	if err := r.db.Write(us); err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *repository) LastId() (int, error) {
	var us []User
	err := r.db.Read(us)
	if err != nil {
		return 0, err
	}
	if len(us) == 0 {
		return 0, nil
	}
	return us[len(us)-1].Id, nil
}

func (r *repository) GetAllUsers() (users []User, err error) {
	err = r.db.Read(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) GetOneUser() {
}

func (r *repository) SaveUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error) {
	u := User{Id, Name, LastName, Email, Age, Height, Active, DateCreate}
	users = append(users, u)
	//lastId = u.Id
	return u, nil
}

func (r *repository) UpdateUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error) {
	u := User{Name: Name, LastName: LastName, Email: Email, Age: Age, Height: Height, Active: Active, DateCreate: DateCreate}
	updated := false
	for i := range users {
		if users[i].Id == Id {
			u.Id = Id
			users[i] = u
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("user %d not found", Id)
	}
	return u, nil
}

func (r *repository) DeleteUser(Id int) error {
	deleted := false
	var index int

	for i := range users {
		if users[i].Id == Id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("user %d not found", Id)
	}

	users = append(users[:index], users[index+1:]...)
	return nil
}

func (r *repository) UpdatePatch(Id int, LastName string, Age int) (User, error) {
	var u User
	updated := false

	for i := range users {
		if users[i].Id == Id {
			users[i].LastName = LastName
			users[i].Age = Age
			updated = true
			u = users[i]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("user %d not found", Id)
	}
	return u, nil
}
