package users

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	dummyData []User
}

func (s StubStore) Read(data interface{}) error {
	users := data.(*[]User)
	*users = s.dummyData
	return nil
}
func (s StubStore) Write(data interface{}) error {
	return nil
}

type MockStore struct {
	ReadWasCalled bool
	dummyData     []User
}

func (m *MockStore) Read(data interface{}) error {
	m.ReadWasCalled = true
	users := data.(*[]User)
	*users = m.dummyData
	return nil
}
func (m MockStore) Write(data interface{}) error {
	return nil
}

func TestReadAll(t *testing.T) {
	lenUsers := 2
	expUsers := []User{
		{
			Id:       2,
			Name:     "Gas2",
			LastName: "Martin2",
			Email:    "gaston2.gmartin@mercadolibre.com",
			Age:      23,
			Height:   165,
			Active:   true,
		},
		{
			Id:       4,
			Name:     "Gas4",
			LastName: "Martin4",
			Email:    "gaston4.gmartin@mercadolibre.com",
			Age:      23,
			Height:   165,
			Active:   true,
		},
	}
	db := StubStore{dummyData: expUsers}
	repo := NewRepository(db)
	users, err := repo.GetAll()

	assert.Nil(t, err)
	assert.True(t, len(users) == lenUsers)
	assert.True(t, reflect.DeepEqual(users, expUsers))
}

func TestUpdateLastNameAndAge(t *testing.T) {
	ExpLastNameUpdated := "After"
	ExpAgeUpdated := 45
	users := []User{
		{
			Id:       25,
			Name:     "gas",
			LastName: "Before",
			Email:    "example@mail.com",
			Age:      23,
			Height:   191,
			Active:   true,
		},
	}
	db := MockStore{ReadWasCalled: false, dummyData: users}
	repo := NewRepository(&db)
	user, err := repo.UpdateLastNameAndAge(25, ExpLastNameUpdated, ExpAgeUpdated)

	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, user.LastName, ExpLastNameUpdated)
	assert.Equal(t, user.Age, ExpAgeUpdated)
}

func TestUpdateLastNameAndAgeBadID(t *testing.T) {

	db := MockStore{ReadWasCalled: false}
	repo := NewRepository(&db)
	user, err := repo.UpdateLastNameAndAge(15, "", 0)

	assert.Error(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, user, User{})
}
