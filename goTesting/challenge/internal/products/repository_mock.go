package products

type mockRepository struct {
	GetInvoked bool
	ErrGet     error
}

func (r *mockRepository) GetAllBySeller(sellerID string) (prodList []Product, err error) {
	r.GetInvoked = true
	if r.ErrGet != nil {
		err = r.ErrGet
		return
	}
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return
}
