package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"encoding/json"
)

func LoadFromJSON[T any](filename string) (T, error) {
	var result T
	data, err := os.ReadFile(filename)
	if err != nil {
		return result, fmt.Errorf("error: utils: LoadFromJSON: ReadFile: %q: %w", filename, err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("error: utils: LoadFromJSON: Unmarshal: %q: %w", filename, err)
	}

	return result, nil
}

func SaveToJSON(filename string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("utils: SaveToJSON: marshal: %w", err)
	}
	data = append(data, '\n')

	dir := filepath.Dir(filename)
	tmp, err := os.CreateTemp(dir, filepath.Base(filename)+".tmp-*")
	if err != nil {
		return fmt.Errorf("utils: SaveToJSON: create temp file in %q: %w", dir, err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("utils: SaveToJSON: write temp file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return fmt.Errorf("utils: SaveToJSON: sync temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("utils: SaveToJSON: close temp file: %w", err)
	}
	if err := os.Chmod(tmpName, 0o644); err != nil {
		return fmt.Errorf("utils: SaveToJSON: chmod temp file: %w", err)
	}
	if err := os.Rename(tmpName, filename); err != nil {
		return fmt.Errorf("utils: SaveToJSON: rename to %q: %w", filename, err)
	}
	return nil
}
