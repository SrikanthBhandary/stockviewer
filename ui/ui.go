// ui/ui.go
package ui

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"stockviewer/finnhub/client"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func RunUI() {
	var content *fyne.Container

	// Initialize a new Fyne application
	myApp := app.New()
	myWindow := myApp.NewWindow("Stock Viewer")

	// Retrieve the Finnhub API key from the environment variable
	apiKey := os.Getenv("API_KEY")
	webClient := client.NewWebClient(apiKey, &http.Client{}) // Use your actual http.Client implementation

	// Create a WebClient instance to interact with the Finnhub API
	symbolEntry := widget.NewEntry()

	// Declare a variable to hold the table widget
	var table *widget.Table

	// Create a button for fetching and displaying stock data
	getButton := widget.NewButton("Get Stock Data", func() {
		// Retrieve the entered stock symbol
		symbol := symbolEntry.Text
		count := 1 // Number of data points to retrieve

		// Retrieve stock data from the Finnhub API
		response, err := webClient.Get(context.Background(), symbol, count)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Remove any existing table from the content container
		content.Remove(table)

		// Create a new table for displaying stock data
		table = widget.NewTable(
			func() (int, int) {
				return len(response.O), 4
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("wide content")
			}, func(cell widget.TableCellID, cellObj fyne.CanvasObject) {
				row := cell.Row
				col := cell.Col

				if row == 0 { // Header row
					switch col {
					case 0:
						cellObj.(*widget.Label).SetText("Open")
					case 1:
						cellObj.(*widget.Label).SetText("High")
					case 2:
						cellObj.(*widget.Label).SetText("Low")
					case 3:
						cellObj.(*widget.Label).SetText("Close")
					}
				} else { // Data rows
					switch col {
					case 0:
						cellObj.(*widget.Label).SetText(fmt.Sprintf("%.2f", response.O[row-1]))
					case 1:
						cellObj.(*widget.Label).SetText(fmt.Sprintf("%.2f", response.H[row-1]))
					case 2:
						cellObj.(*widget.Label).SetText(fmt.Sprintf("%.2f", response.L[row-1]))
					case 3:
						cellObj.(*widget.Label).SetText(fmt.Sprintf("%.2f", response.C[row-1]))
					}
				}
			},
		)
		// Add the table to the content container
		content.Add(table)
	})

	// Create a horizontal container for UI elements
	horizontaContainer := container.New(layout.NewFormLayout(), widget.NewLabel("Enter Symbol:"),
		symbolEntry)

	// Create a vertical container to hold the entire UI layout
	content = container.New(
		layout.NewGridLayoutWithRows(3),
		horizontaContainer,
		getButton,
	)

	// Set the content of the application window
	myWindow.SetContent(content)

	// Set the window size
	myWindow.Resize(fyne.NewSize(400, 250))

	// Show and run the application
	myWindow.ShowAndRun()
}
