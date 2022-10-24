package users

import (
	"fmt"
	"testing"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class4-tt/internal/domain"
	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class4-tt/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryUpdate(t *testing.T) {
	// Arrange.
	initialData := []domain.User{
		{
			Id:       5,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   191,
			Active:   true,
		},
	}
	expUser := domain.User{
		Id:       5,
		Name:     "test",
		LastName: "test",
		Email:    "test@mail.com",
		Age:      32,
		Height:   151,
		Active:   true,
	}
	db := store.MockStore{ReadWasCalled: false, DummyData: initialData}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.Update(5, expUser.Name, expUser.LastName, expUser.Email, expUser.Age, expUser.Height, expUser.Active)

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Equal(t, expUser, result)
}
func TestRepositoryUpdateBadId(t *testing.T) {
	// Arrange.
	invalidID := 21
	ErrNotFound := fmt.Errorf("user with id %d not found", invalidID)
	initialData := []domain.User{
		{
			Id:       5,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   191,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.Update(invalidID, "", "", "", 0, 0, false)

	// Assert.
	assert.EqualError(t, ErrNotFound, err.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Empty(t, result)
}

func TestRepositoryDelete(t *testing.T) {
	// Arrange.
	clientIDToDelete := 1
	initialData := []domain.User{
		{
			Id:       1,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   191,
			Active:   true,
		},
		{
			Id:       2,
			Name:     "test",
			LastName: "test",
			Email:    "test@mail.com",
			Age:      34,
			Height:   163,
			Active:   true,
		},
	}
	expdata := []domain.User{
		{
			Id:       2,
			Name:     "test",
			LastName: "test",
			Email:    "test@mail.com",
			Age:      34,
			Height:   163,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData}
	repo := NewRepository(&db)
	// Act.
	err := repo.Delete(clientIDToDelete)
	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Equal(t, expdata, db.DummyData)
}

func TestRepositoryDeleteBadId(t *testing.T) {
	// Arrange.
	invalidID := 4
	ErrNotFound := fmt.Errorf("user with id %d not found", invalidID)
	initialData := []domain.User{
		{
			Id:       1,
			Name:     "Gaston",
			LastName: "Maritn",
			Email:    "test@mail.com",
			Age:      23,
			Height:   191,
			Active:   true,
		},
		{
			Id:       2,
			Name:     "test",
			LastName: "test",
			Email:    "test@mail.com",
			Age:      34,
			Height:   163,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData}
	repo := NewRepository(&db)
	// Act.
	err := repo.Delete(invalidID)
	// Assert.
	assert.EqualError(t, err, ErrNotFound.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Equal(t, initialData, db.DummyData)
}
