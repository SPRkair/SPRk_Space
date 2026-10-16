package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvFile(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return fmt.Errorf("utils: LoadFromEnv: Load: %s: %w", filename, err)
	}
	return nil
}
