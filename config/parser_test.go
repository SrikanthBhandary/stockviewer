// Package config provides functionality to parse configuration files in YAML format.
// It includes a function to parse a YAML file containing a list of symbols.
package config

import (
	"os"
	"testing"
)

// TestParseSymbolsFile is a test case for the ParseSymbolsFile function.
// It creates a temporary test YAML file, writes test content to it, and parses the symbols using ParseSymbolsFile.
// It then compares the parsed symbols with the expected symbols to verify the parsing functionality.
func TestParseSymbolsFile(t *testing.T) {
	// Create a temporary test YAML file
	yamlContent := `
symbols:
  - HPQ
  - IBM
  - AAPL
  - MSFT
`
	tempFile, err := os.CreateTemp("", "test-symbols.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write the test YAML content to the temporary file
	_, err = tempFile.WriteString(yamlContent)
	if err != nil {
		t.Fatalf("Failed to write test YAML content: %v", err)
	}
	tempFile.Close()

	// Test the ParseSymbolsFile function
	symbols, err := ParseSymbolsFile(tempFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Define the expected symbols
	expectedSymbols := []string{"HPQ", "IBM", "AAPL", "MSFT"}

	// Compare the parsed symbols with the expected symbols
	if len(symbols) != len(expectedSymbols) {
		t.Errorf("Number of symbols does not match: got %d, want %d", len(symbols), len(expectedSymbols))
	}

	for i := range symbols {
		if symbols[i] != expectedSymbols[i] {
			t.Errorf("Symbol mismatch at index %d: got %s, want %s", i, symbols[i], expectedSymbols[i])
		}
	}
}
