package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadFromYaml(filename string, structure interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromYaml: ReadFile:  %w", err)
	}

	err = yaml.Unmarshal(data, structure)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromYaml: Unmarshal:  %w", err)
	}

	return nil
}

func SaveToYaml(filename string, structure interface{}) error {
	data, err := yaml.Marshal(structure)
	if err != nil {
		return fmt.Errorf("error: utils: SaveToYaml: Marshal: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error: utils: SaveToYaml: WriteFile: %w", err)
	}

	return nil
}

func AppendToYaml(filename string, structure interface{}) error {
	var existing []interface{}
	if err := LoadFromYaml(filename, &existing); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("error: utils: AppendToYaml: Load: %q: %w", filename, err)
		}
	}
	existing = append(existing, structure)
	if err := SaveToYaml(filename, existing); err != nil {
		return fmt.Errorf("error: utils: AppendToYaml: Save: %q: %w", filename, err)
	}
	return nil
}
