package utils

import (
	"fmt"
	"os"
)

func LoadFromTxt(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("read file %q: %w", filename, err)
	}
	return string(data), nil
}

func SaveToTxt(filename string, text string) error {
	if err := os.WriteFile(filename, []byte(text), 0644); err != nil {
		return fmt.Errorf("write file %q: %w", filename, err)
	}
	return nil
}

func AppendToTxt(filename string, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file %q: %w", filename, err)
	}

	if _, err := f.WriteString(text); err != nil {
		f.Close()
		return fmt.Errorf("append to file %q: %w", filename, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("close file %q: %w", filename, err)
	}
	return nil
}
