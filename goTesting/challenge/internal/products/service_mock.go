package products

type mockService struct {
	repository Repository
	GetInvoked bool
	ErrGet     error
}

func (s *mockService) GetAllBySeller(sellerID string) ([]Product, error) {
	if s.ErrGet != nil {
		return nil, s.ErrGet
	}
	data, err := s.repository.GetAllBySeller(sellerID)
	if err != nil {
		return nil, err
	}
	return data, err
}
