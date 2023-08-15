package model

// Response represents the stock data response from the Finnhub API.
type Response struct {
	C []float64 `json:"c"` // Closing prices for each data point.
	H []float64 `json:"h"` // High prices for each data point.
	L []float64 `json:"l"` // Low prices for each data point.
	O []float64 `json:"o"` // Opening prices for each data point.
	S string    `json:"s"` // Symbol of the stock.
	T []int     `json:"t"` // Timestamps for each data point.
	V []int     `json:"v"` // Volume of shares traded for each data point.
}
