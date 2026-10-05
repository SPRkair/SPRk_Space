package utils

import (
	"fmt"
	"os"

	"encoding/json"
)

func LoadFromJSON(filename string, structure interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromJSON: ReadFile: %w", err)
	}

	err = json.Unmarshal(data, structure)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromJSON: Unmarshal: %w", err)
	}

	return nil
}

func SaveToJSON(filename string, structure interface{}) error {
	data, err := json.MarshalIndent(structure, "", "  ")
	if err != nil {
		return fmt.Errorf("error: utils: SaveToJSON: Marshal: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error: utils: SaveToJSON: WriteFile: %w", err)
	}

	return nil
}
