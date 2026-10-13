package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadFromJSON(filename string, structure any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromJSON: ReadFile: %q: %w", filename, err)
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return nil
	}

	err = json.Unmarshal(data, structure)
	if err != nil {
		return fmt.Errorf("error: utils: LoadFromJSON: Unmarshal: %q: %w", filename, err)
	}

	return nil
}

func SaveToJSON(filename string, structure any) error {
	data, err := json.MarshalIndent(structure, "", "  ")
	if err != nil {
		return fmt.Errorf("error: utils: SaveToJSON: Marshal: %q: %w", filename, err)
	}
	data = append(data, '\n')

	dir := filepath.Dir(filename)
	tmp, err := os.CreateTemp(dir, filepath.Base(filename)+".tmp-*")
	if err != nil {
		return fmt.Errorf("error: utils: SaveToJSON: CreateTemp: %q: %w", dir, err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("error: utils: SaveToJSON: Write: %q: %w", tmpName, err)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return fmt.Errorf("error: utils: SaveToJSON: Sync: %q: %w", tmpName, err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("error: utils: SaveToJSON: Close: %q: %w", tmpName, err)
	}
	if err := os.Chmod(tmpName, 0o644); err != nil {
		return fmt.Errorf("error: utils: SaveToJSON: Chmod: %q: %w", tmpName, err)
	}
	if err := os.Rename(tmpName, filename); err != nil {
		return fmt.Errorf("error: utils: SaveToJSON: Rename: %q: %w", filename, err)
	}

	return nil
}
