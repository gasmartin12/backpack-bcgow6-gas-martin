package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	data []Product
}

func (r *mockRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product

	for _, v := range r.data {
		if sellerID == v.SellerID {
			prodList = append(prodList, v)
		}
	}

	if len(prodList) == 0 {
		return []Product{}, errors.New("not found")
	}

	return prodList, nil
}

func TestGetAllBySeller(t *testing.T) {

	expected := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}

	repo := mockRepository{
		data: []Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		},
	}

	service := NewService(&repo)

	out, err := service.GetAllBySeller("FEX112AC")

	assert.Nil(t, err)
	assert.Equal(t, expected, out)
}

func TestGetAllBySellerWithError(t *testing.T){

	expected := errors.New("not found")

	repo := mockRepository{
		data: []Product{},
	}

	service := NewService(&repo)

	out, err := service.GetAllBySeller("FEX112AC")

	assert.Empty(t, out)
	assert.Equal(t, expected, err)
}