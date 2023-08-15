// Package client provides a WebClient for interacting with the Finnhub API to retrieve stock data.
package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"stockviewer/finnhub/model"
)

// HttpClient defines the interface for making HTTP requests.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const FINHUBURL = "https://finnhub.io"

var errBadRequest = errors.New("invalid response from finhub")

// WebClient is a client for interacting with the Finnhub API.
type WebClient struct {
	apiKey     string
	httpClient HttpClient
}

// NewWebClient creates a new WebClient instance.
func NewWebClient(apiKey string, httpClient HttpClient) *WebClient {
	return &WebClient{apiKey: apiKey, httpClient: httpClient}
}

// Get retrieves stock data from the Finnhub API for a given symbol and count.
func (w *WebClient) Get(ctx context.Context, symbol string, count int) (*model.Response, error) {
	url := fmt.Sprintf("%s/api/v1/stock/candle?symbol=%s&resolution=D&count=%d&token=%s",
		FINHUBURL, symbol, count, w.apiKey)
	fmt.Println(url)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	response, err := w.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var stockResponse model.Response
	if response.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&stockResponse)
		if err != nil {
			return nil, err
		}
		return &stockResponse, nil
	}
	return nil, errBadRequest
}
