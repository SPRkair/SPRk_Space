package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadFromEnv(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromEnv: Load:  %w", err)
	}
	return nil
}
