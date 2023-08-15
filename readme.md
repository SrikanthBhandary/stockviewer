---

## Stock Viewer Application using Fyne

The Stock Viewer Application is a graphical user interface (GUI) application that allows users to retrieve and view stock data using the Finnhub API. It provides a simple way to enter a stock symbol and fetch the corresponding stock data, which is then displayed in a table.

### Features:

1. **Symbol Input**: Users can enter a stock symbol in the provided input field.

2. **Data Retrieval**: Clicking the "Get Stock Data" button fetches stock data from the Finnhub API based on the entered stock symbol.

3. **Data Display**: Retrieved stock data is displayed in a table with columns representing the Open, High, Low, and Close values for each data point.

### Prerequisites:

1. Install the Fyne package by running the following command:

   ```bash
   go get fyne.io/fyne/v2
   ```

2. Import necessary packages:

   ```go
   import (
       "context"
       "fmt"
       "net/http"
       "os"
       "fyne.io/fyne/v2"
       "fyne.io/fyne/v2/app"
       "fyne.io/fyne/v2/container"
       "fyne.io/fyne/v2/layout"
       "fyne.io/fyne/v2/widget"
       "stockviewer/finnhub/client"
   )
   ```

### Code Explanation:

1. **Initializing the Application**:

   The application is initialized using `app.New()` to create a new Fyne app instance and `NewWindow()` to create the main application window.

2. **API Key and WebClient**:

   The Finnhub API key is stored in the `apiKey` variable, and a `WebClient` instance is created using the provided API key and an HTTP client.

3. **User Interface Elements**:

   - `symbolEntry`: An entry widget that allows users to input the stock symbol.
   - `getButton`: A button that, when clicked, retrieves and displays stock data based on the entered stock symbol.
   - `table`: A table widget for displaying the retrieved stock data.

4. **Button Click Handler** (`getButton`):

   When the "Get Stock Data" button is clicked, the entered stock symbol is used to retrieve stock data from the Finnhub API. The retrieved data is then displayed in the table.

5. **Table Creation and Population**:

   The table is created using `widget.NewTable()` and populated with stock data. The headers for the table columns are added in the first row.

6. **UI Layout**:

   The user interface is organized using the `layout.NewGridLayoutWithRows()` and `layout.NewFormLayout()` layouts. The elements are placed in appropriate containers for alignment and spacing.

7. **Application Window Setup**:

   The application window's content is set using `myWindow.SetContent()`.

8. **Running the Application**:

   The application is displayed using `myWindow.ShowAndRun()`.

### Conclusion:

The Stock Viewer Application provides a simple and user-friendly way to retrieve and view stock data using the Finnhub API. Users can input a stock symbol, click a button to fetch data, and visualize the retrieved data in a table. The application's basic functionality can serve as a foundation for building more advanced stock data visualization tools.

---

Feel free to further enhance the application by adding error handling, additional features, and improved user interfaces.