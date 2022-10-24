package store

import "github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class5-tm/internal/domain"

type MockStore struct {
	ReadWasCalled  bool
	WriteWasCalled bool
	ErrRead        error
	ErrWrite       error
	DummyData      []domain.User
}

func (m *MockStore) Read(data interface{}) error {
	m.ReadWasCalled = true
	if m.ErrRead != nil {
		return m.ErrRead
	}
	*data.(*[]domain.User) = m.DummyData
	return nil
}
func (m *MockStore) Write(data interface{}) error {
	m.WriteWasCalled = true
	if m.ErrWrite != nil {
		return m.ErrWrite
	}
	m.DummyData = data.([]domain.User)
	return nil
}
