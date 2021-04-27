package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDataStore struct{}

func (ds mockDataStore) PackSizesForProduct(productId string) ([]int, bool, error) {
	return []int{250, 500, 1000, 2000, 5000}, true, nil
}

type mockResponse struct {
	Pack250  int `json:"250"`
	Pack500  int `json:"500"`
	Pack1000 int `json:"1000"`
	Pack2000 int `json:"2000"`
	Pack5000 int `json:"5000"`
}

type testCase struct {
	input          string
	expectedRes    mockResponse
	expectedStatus int
}

func TestCalculatePacksToSend(t *testing.T) {
	var tests = []testCase{
		{input: `{
			"productId": "123",
			"itemsOrder": 1
		}`, expectedRes: mockResponse{
			Pack250: 1,
		}, expectedStatus: 200},
		{input: `{
			"productId": "123",
			"itemsOrder": 250
		}`, expectedRes: mockResponse{
			Pack250: 1,
		}, expectedStatus: 200},
		{input: `{
			"productId": "123",
			"itemsOrder": 251
		}`, expectedRes: mockResponse{
			Pack500: 1,
		}, expectedStatus: 200},
		{input: `{
			"productId": "123",
			"itemsOrder": 501
		}`, expectedRes: mockResponse{
			Pack250: 1,
			Pack500: 1,
		}, expectedStatus: 200},
		{input: `{
			"productId": "123",
			"itemsOrder": 12001
		}`, expectedRes: mockResponse{
			Pack250:  1,
			Pack2000: 1,
			Pack5000: 2,
		}, expectedStatus: 200},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.input))
		rec := httptest.NewRecorder()

		ph := Product{dataStore: mockDataStore{}}
		ph.CalculatePacksToSend(rec, req)

		actualRes := mockResponse{}
		json.NewDecoder(rec.Body).Decode(&actualRes)

		assert.Equal(t, tc.expectedStatus, rec.Code)
		assert.Equal(t, tc.expectedRes, actualRes)
	}
}
