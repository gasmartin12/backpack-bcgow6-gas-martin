package users

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class5-tm/internal/domain"
	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class5-tm/pkg/store"
	"github.com/stretchr/testify/assert"
)

// Get all testing
func TestRepositoryGetall(t *testing.T) {
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

	// Act.
	result, err := repo.GetAll()

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, initialData, result)
}
func TestRepositoryGetallReadingError(t *testing.T) {
	// Arrange.
	ErrReadConnection := errors.New("bad connection to database")
	db := store.MockStore{ErrRead: ErrReadConnection}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.GetAll()

	// Assert.
	assert.EqualError(t, err, ErrReadConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.Empty(t, result)
}

// Store testing
func TestRepositoryStore(t *testing.T) {
	// Arrange.
	expUser := domain.User{
		Id:           5,
		Name:         "Gaston",
		LastName:     "Maritn",
		Email:        "test@mail.com",
		Age:          23,
		Height:       166,
		Active:       true,
		CreationDate: time.Time{},
	}
	db := store.MockStore{}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.Store(5, expUser.Name, expUser.LastName, expUser.Email, expUser.Age, expUser.Height, expUser.Active, expUser.CreationDate)

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Equal(t, expUser, result)
}
func TestRepositoryStoreReadingError(t *testing.T) {
	// Arrange.
	ErrReadConnection := errors.New("bad connection to database")
	db := store.MockStore{ErrRead: ErrReadConnection}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.Store(0, "", "", "", 0, 0, false, time.Time{})

	// Assert.
	assert.EqualError(t, err, ErrReadConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Empty(t, result)
}
func TestRepositoryStorewritindingError(t *testing.T) {
	// Arrange.
	ErrWriteConnection := errors.New("bad connection to database")
	db := store.MockStore{ErrWrite: ErrWriteConnection}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.Store(5, "", "", "", 0, 0, false, time.Time{})

	// Assert.
	assert.EqualError(t, err, ErrWriteConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Empty(t, result)
}

// Update testing
func TestRepositoryUpdate(t *testing.T) {
	// Arrange.
	initialData := []domain.User{
		{
			Id:       5,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
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
			Height:   166,
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
func TestRepositoryUpdateReadingError(t *testing.T) {
	// Arrange.
	ErrReadConnection := errors.New("bad connection to database")
	db := store.MockStore{ErrRead: ErrReadConnection}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.Update(0, "", "", "", 0, 0, false)

	// Assert.
	assert.EqualError(t, err, ErrReadConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Empty(t, result)
}
func TestRepositoryUpdatewritindingError(t *testing.T) {
	// Arrange.
	ErrWriteConnection := errors.New("bad connection to database")
	initialData := []domain.User{
		{
			Id:       5,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData, ErrWrite: ErrWriteConnection}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.Update(5, "", "", "", 0, 0, false)

	// Assert.
	assert.EqualError(t, err, ErrWriteConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Empty(t, result)
}

// Update last name and age testing
func TestRepositoryUpdateLastNameAndAge(t *testing.T) {
	// Arrange.
	initialData := []domain.User{
		{
			Id:       5,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
	}
	expUser := domain.User{
		Id:       5,
		Name:     "test",
		LastName: "test",
		Email:    "test@mail.com",
		Age:      43,
		Height:   191,
		Active:   true,
	}
	db := store.MockStore{DummyData: initialData}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.UpdateLastNameAndAge(5, expUser.LastName, expUser.Age)

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Equal(t, expUser, result)
}
func TestRepositoryUpdateLastNameAndAgeBadId(t *testing.T) {
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
			Height:   166,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.UpdateLastNameAndAge(invalidID, "", 0)

	// Assert.
	assert.EqualError(t, ErrNotFound, err.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Empty(t, result)
}
func TestRepositoryUpdateLastNameAndAgeReadingError(t *testing.T) {
	// Arrange.
	ErrReadConnection := errors.New("bad connection to database")
	db := store.MockStore{ErrRead: ErrReadConnection}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.UpdateLastNameAndAge(0, "", 0)

	// Assert.
	assert.EqualError(t, err, ErrReadConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Empty(t, result)
}
func TestRepositoryUpdateLastNameAndAgewritindingError(t *testing.T) {
	// Arrange.
	ErrWriteConnection := errors.New("bad connection to database")
	initialData := []domain.User{
		{
			Id:       5,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData, ErrWrite: ErrWriteConnection}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.UpdateLastNameAndAge(5, "", 0)

	// Assert.
	assert.EqualError(t, err, ErrWriteConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.True(t, db.WriteWasCalled)
	assert.Empty(t, result)
}

// Delete testing
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
	// Act.
	err := repo.Delete(invalidID)
	// Assert.
	assert.EqualError(t, err, ErrNotFound.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
	assert.Equal(t, initialData, db.DummyData)
}
func TestRepositorydeleteReadingError(t *testing.T) {
	// Arrange.
	ErrReadConnection := errors.New("bad connection to database")
	db := store.MockStore{ErrRead: ErrReadConnection}
	repo := NewRepository(&db)

	// Act.
	err := repo.Delete(0)

	// Assert.
	assert.EqualError(t, err, ErrReadConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.False(t, db.WriteWasCalled)
}
func TestRepositorydeletewritindingError(t *testing.T) {
	// Arrange.
	ErrWriteConnection := errors.New("bad connection to database")
	initialData := []domain.User{
		{
			Id:       5,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
	}
	db := store.MockStore{DummyData: initialData, ErrWrite: ErrWriteConnection}
	repo := NewRepository(&db)

	// Act.
	err := repo.Delete(5)

	// Assert.
	assert.EqualError(t, err, ErrWriteConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, initialData, db.DummyData)
	assert.True(t, db.WriteWasCalled)
}

// Last id testing
func TestRepositoryLastId(t *testing.T) {
	// Arrange.
	expID := 19
	initialData := []domain.User{
		{
			Id:       18,
			Name:     "Gaston",
			LastName: "Martin",
			Email:    "test@mail.com",
			Age:      23,
			Height:   166,
			Active:   true,
		},
		{
			Id:       19,
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
	result, err := repo.LastId()

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, expID, result)
}
func TestRepositoryLastIdEmptyInitalUserData(t *testing.T) {
	// Arrange.
	expID := 0
	db := store.MockStore{}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.LastId()

	// Assert.
	assert.Nil(t, err)
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, expID, result)
}
func TestRepositoryLastIdReadingError(t *testing.T) {
	// Arrange.
	expID := 0
	ErrReadConnection := errors.New("bad connection to database")
	db := store.MockStore{ErrRead: ErrReadConnection}
	repo := NewRepository(&db)

	// Act.
	result, err := repo.LastId()

	// Assert.
	assert.EqualError(t, err, ErrReadConnection.Error())
	assert.True(t, db.ReadWasCalled)
	assert.Equal(t, expID, result)
}
