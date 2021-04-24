package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockJsonDataStore struct{}

func (ds mockJsonDataStore) PackSizesForProduct(productId string) ([]int, bool, error) {
	return []int{250, 500, 1000, 2000, 5000}, true, nil
}

type mockResponse struct {
	Pack500 int `json:"500"`
}

func TestCalculatePacksToSend(t *testing.T) {
	body := `{
		"productId": "123",
		"itemsOrder": 251
	}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	rec := httptest.NewRecorder()

	ph := ProductHandler{dataStore: mockJsonDataStore{}}
	ph.CalculatePacksToSend(rec, req)

	expectedRes := mockResponse{
		Pack500: 1,
	}
	actualRes := mockResponse{}
	json.NewDecoder(rec.Body).Decode(&actualRes)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expectedRes, actualRes)
}
