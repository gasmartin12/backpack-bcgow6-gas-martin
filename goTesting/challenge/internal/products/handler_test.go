package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	BaseURL = "/api/v1/products"
)

func createServer(err error) (engine *gin.Engine) {
	// Initialization of services
	productsRepository := NewRepository()
	productsService := mockService{
		repository: productsRepository,
		ErrGet:     err,
	}
	productsHandler := NewHandler(&productsService)
	engine = gin.New()
	// Endpoints
	// Products
	productsPrefix := engine.Group(BaseURL)
	{ // -> This does nothing, but helps with clean code
		// GETs
		productsPrefix.GET("", productsHandler.GetProducts)
	}
	return
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	request.Header.Add("Content-Type", "application/json")

	return request, httptest.NewRecorder()
}

func testGetProductsOk(t *testing.T) {
	server := createServer(nil)
	getSellerID := "?seller_id=FEX112AC"
	expected := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	request, responseRecorder := createRequestTest(http.MethodGet, BaseURL+getSellerID, "")
	server.ServeHTTP(responseRecorder, request)
	var responseObject []Product
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseObject)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Nil(t, err)
	assert.Equal(t, expected, responseObject)
}

func testGetProductsBadSellerID(t *testing.T) {
	server := createServer(nil)
	expected := map[string]string{
		"error": "seller_id query param is required",
	}
	request, responseRecorder := createRequestTest(http.MethodGet, BaseURL, "")
	server.ServeHTTP(responseRecorder, request)
	var responseObject map[string]string
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseObject)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Nil(t, err)
	assert.Equal(t, expected, responseObject)
}

func testGetProductsWithErr(t *testing.T) {
	getSellerID := "?seller_id=FEX112AC"
	expectedErr := errors.New("forced error in service")
	server := createServer(expectedErr)
	expected := map[string]string{
		"error": expectedErr.Error(),
	}
	request, responseRecorder := createRequestTest(http.MethodGet, BaseURL+getSellerID, "")
	server.ServeHTTP(responseRecorder, request)
	var responseObject map[string]string
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &responseObject)
	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	assert.Nil(t, err)
	assert.Equal(t, expected, responseObject)
}

func testBadPathSuit(t *testing.T) {
	t.Run("Bad SellerID", testGetProductsBadSellerID)
	t.Run("Error on service", testGetProductsWithErr)
}

func TestSuitGET(t *testing.T) {
	t.Run("Happy path", testGetProductsOk)
	t.Run("Sad Path", testBadPathSuit)
}
