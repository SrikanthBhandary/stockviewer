// Package client provides a WebClient for interacting with the Finnhub API to retrieve stock data.
package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"stockviewer/finnhub/model"
	"testing"
)

// MockHttpClient is a mock implementation of the HttpClient interface for testing purposes.
type MockHttpClient struct{}

// Do is the mock implementation of the Do method of the HttpClient interface.
func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	// Create a mock HTTP response for testing
	response := httptest.NewRecorder()
	response.WriteString(`{"O":[100,101,102],"H":[110,111,112],"L":[90,91,92],"C":[105,106,107]}`)

	return response.Result(), nil
}

func TestWebClient_Get(t *testing.T) {
	// Create a mock HTTP client
	mockClient := &MockHttpClient{}

	// Create a mock context for testing
	ctx := context.Background()

	// Create a new WebClient instance with the mock client
	webClient := NewWebClient("your_api_key", mockClient)

	// Call the Get method to retrieve stock data
	symbol := "AAPL"
	count := 3
	response, err := webClient.Get(ctx, symbol, count)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Define the expected response
	expectedResponse := &model.Response{
		O: []float64{100, 101, 102},
		H: []float64{110, 111, 112},
		L: []float64{90, 91, 92},
		C: []float64{105, 106, 107},
	}

	// Compare the response with the expected response
	if !compareResponses(response, expectedResponse) {
		t.Errorf("Response does not match expected: got %+v, want %+v", response, expectedResponse)
	}
}

func compareResponses(a, b *model.Response) bool {
	// Implement your own logic to compare the response structures
	// For this example, we're just comparing the slices for simplicity
	return compareSlices(a.O, b.O) &&
		compareSlices(a.H, b.H) &&
		compareSlices(a.L, b.L) &&
		compareSlices(a.C, b.C)
}

func compareSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
