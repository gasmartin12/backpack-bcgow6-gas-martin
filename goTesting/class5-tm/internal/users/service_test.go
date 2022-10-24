package users

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class5-tm/internal/domain"
	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class5-tm/pkg/store"
	"github.com/stretchr/testify/assert"
)

// Get all testing
func TestServiceGetall(t *testing.T) {
	// Arrange.
	initialData := []domain.User{
		{
			Id:       1,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
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
	service := NewService(repo)

	// Act.
	result, err := service.GetAll()

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, initialData, result)
}
func TestServiceGetallReadingError(t *testing.T) {
	// Arrange.
	ErrReadConnection := errors.New("bad connection to database")
	db := store.MockStore{ErrRead: ErrReadConnection}
	repo := NewRepository(&db)
	service := NewService(repo)

	// Act.
	result, err := service.GetAll()

	// Assert.
	assert.EqualError(t, err, ErrReadConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.Empty(t, result)
}

// Store testing
func TestServiceStore(t *testing.T) {
	// Arrange.
	expUser := domain.User{
		Id:       1,
		Name:     "Gaston",
		LastName: "Martin",
		Email:    "test@mail.com",
		Age:      23,
		Height:   166,
		Active:   true,
	}
	db := store.MockStore{}
	repo := NewRepository(&db)
	service := NewService(repo)

	// Act.
	result, err := service.Store(expUser.Name, expUser.LastName, expUser.Email, expUser.Age, expUser.Height)

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Equal(t, expUser.Id, result.Id)
	assert.Equal(t, expUser.Name, result.Name)
	assert.Equal(t, expUser.LastName, result.LastName)
	assert.Equal(t, expUser.Email, result.Email)
	assert.Equal(t, expUser.Age, result.Age)
	assert.Equal(t, expUser.Height, result.Height)
	assert.Equal(t, expUser.Active, result.Active)
}
func TestServiceStoreReadingErrorByGetLastId(t *testing.T) {
	// Arrange.
	ErrReadConnection := errors.New("bad connection to database")
	ErrReadToGiveId := errors.New("id could not be generated")
	db := store.MockStore{ErrRead: ErrReadConnection}
	repo := NewRepository(&db)
	service := NewService(repo)

	// Act.
	result, err := service.Store("", "", "", 0, 0)

	// Assert.
	assert.EqualError(t, err, ErrReadToGiveId.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Empty(t, result)
}
func TestServiceStoreWritingError(t *testing.T) {
	// Arrange.
	ErrWriteConnection := errors.New("bad connection to database")
	ErrCreateUser := errors.New("user could not be created")
	db := store.MockStore{ErrWrite: ErrWriteConnection}
	repo := NewRepository(&db)
	service := NewService(repo)

	// Act.
	result, err := service.Store("", "", "", 0, 0)

	// Assert.
	assert.EqualError(t, err, ErrCreateUser.Error())
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Empty(t, result)
}

// Update testing
func TestServiceUpdate(t *testing.T) {
	// Arrange.
	initialData := []domain.User{
		{
			Id:       25,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
	}
	expUser := domain.User{
		Id:       25,
		Name:     "test",
		LastName: "test",
		Email:    "test@mail.com",
		Age:      32,
		Height:   151,
		Active:   true,
	}
	db := store.MockStore{ReadWasCalled: false, DummyData: initialData}
	repo := NewRepository(&db)
	service := NewService(repo)

	// Act.
	result, err := service.Update(25, expUser.Name, expUser.LastName, expUser.Email, expUser.Age, expUser.Height, expUser.Active)

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Equal(t, expUser, result)
}
func TestServiceUpdateBadId(t *testing.T) {
	// Arrange.
	invalidID := 21
	ErrNotFound := fmt.Errorf("user with id %d not found", invalidID)
	initialData := []domain.User{
		{
			Id:       25,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData}
	repo := NewRepository(&db)
	service := NewService(repo)

	// Act.
	result, err := service.Update(invalidID, "", "", "", 0, 0, false)

	// Assert.
	assert.EqualError(t, ErrNotFound, err.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Empty(t, result)
}

// Update last name and age testing
func TestServiceUpdateLastNameAndAge(t *testing.T) {
	// Arrange.
	initialData := []domain.User{
		{
			Id:       25,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
	}
	expUser := domain.User{
		Id:       25,
		Name:     "Gaston",
		LastName: "Martin",
		Email:    "test@mail.com",
		Age:      43,
		Height:   166,
		Active:   true,
	}
	db := store.MockStore{DummyData: initialData}
	repo := NewRepository(&db)
	service := NewService(repo)

	// Act.
	result, err := service.UpdateLastNameAndAge(25, expUser.LastName, expUser.Age)

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Equal(t, expUser, result)
}
func TestServiceUpdateLastNameAndAgeBadId(t *testing.T) {
	// Arrange.
	invalidID := 21
	ErrNotFound := fmt.Errorf("user with id %d not found", invalidID)
	initialData := []domain.User{
		{
			Id:       25,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData}
	repo := NewRepository(&db)
	service := NewService(repo)

	// Act.
	result, err := service.UpdateLastNameAndAge(invalidID, "", 0)

	// Assert.
	assert.EqualError(t, ErrNotFound, err.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Empty(t, result)
}

// Delete testing
func TestServiceDelete(t *testing.T) {
	// Arrange.
	clientIDToDelete := 1
	initialData := []domain.User{
		{
			Id:       1,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
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
	service := NewService(repo)
	// Act.
	err := service.Delete(clientIDToDelete)
	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Equal(t, expdata, db.DummyData)
}
func TestServiceDeleteBadId(t *testing.T) {
	// Arrange.
	invalidID := 4
	ErrNotFound := fmt.Errorf("user with id %d not found", invalidID)
	initialData := []domain.User{
		{
			Id:       1,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
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
	service := NewService(repo)

	// Act.
	err := service.Delete(invalidID)

	// Assert.
	assert.EqualError(t, err, ErrNotFound.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Equal(t, initialData, db.DummyData)
}
