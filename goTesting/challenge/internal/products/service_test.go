package products

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testServiceGetAllBySellerOk(t *testing.T) {
	getSellerID := "FEX112AC"
	expected := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	mockProductRepository := mockRepository{}
	productService := NewService(&mockProductRepository)
	data, err := productService.GetAllBySeller(getSellerID)
	assert.Nil(t, err)
	assert.Equal(t, expected, data)
	assert.True(t, mockProductRepository.GetInvoked)
}

func testServiceGetAllBySellerErr(t *testing.T) {
	getSellerID := "FEX112AC"
	expectedErr := errors.New("forced error in repository")
	var expected []Product
	mockProductRepository := mockRepository{ErrGet: expectedErr}
	productService := NewService(&mockProductRepository)
	data, err := productService.GetAllBySeller(getSellerID)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Equal(t, expected, data)
	assert.True(t, mockProductRepository.GetInvoked)
}

func TestServiceGetAllBySeller(t *testing.T) {
	t.Run("Happy Path", testServiceGetAllBySellerOk)
	t.Run("Sad Path", testServiceGetAllBySellerErr)
}
