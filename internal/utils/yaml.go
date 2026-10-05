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
