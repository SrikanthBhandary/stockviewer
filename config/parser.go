// Package config provides functionality to parse configuration files in YAML format.
// It includes a function to parse a YAML file containing a list of symbols.
package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// SymbolsConfig represents the structure of the symbols configuration.
// It contains a list of symbols read from the YAML file.
type SymbolsConfig struct {
	Symbols []string `yaml:"symbols"`
}

// ParseSymbolsFile reads a YAML file containing a list of symbols and parses it.
// It returns the list of symbols present in the file.
func ParseSymbolsFile(filename string) ([]string, error) {
	// Read the YAML file
	yamlData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse the YAML data into the SymbolsConfig struct
	var symbolsConfig SymbolsConfig
	err = yaml.Unmarshal(yamlData, &symbolsConfig)
	if err != nil {
		return nil, err
	}

	return symbolsConfig.Symbols, nil
}
