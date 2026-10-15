package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func LoadFromIni(filename string, structure interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromIni: ReadFile:  %w", err)
	}

	cfg, err := ini.Load(data)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromIni: Load: %w", err)
	}
	err = cfg.MapTo(structure)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromIni: MapTo: %w", err)
	}
	return nil
}

func SaveToIni(filename string, structure interface{}) error {
	cfg := ini.Empty()

	err := cfg.ReflectFrom(structure)
	if err != nil {
		return fmt.Errorf("error: utils: SaveToIni: ReflectFrom: %w", err)
	}

	err = cfg.SaveTo(filename)
	if err != nil {
		return fmt.Errorf("error: utils: SaveToIni: SaveTo: %w", err)
	}

	return nil
}

func AppendToIni(filename string, structure interface{}) error {
	cfg, err := ini.Load(filename)
	if err != nil {
		return fmt.Errorf("error: utils: AppendToIni: Load: %q: %w", filename, err)
	}

	err = cfg.ReflectFrom(structure)
	if err != nil {
		return fmt.Errorf("error: utils: AppendToIni: ReflectFrom: %q: %w", filename, err)
	}

	err = cfg.SaveTo(filename)
	if err != nil {
		return fmt.Errorf("error: utils: AppendToIni: SaveTo: %q: %w", filename, err)
	}

	return nil
}
