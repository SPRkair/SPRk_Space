package utils

import (
	"fmt"
	"path/filepath"
)

func LoadFrom[T any](filename string, structure *T) error {
	switch filepath.Ext(filename) {
	case ".yaml":
		return LoadFromYaml(filename, structure)
	case ".json":
		return LoadFromJSON(filename, structure)
	case ".ini":
		return LoadFromIni(filename, structure)
	default:
		return fmt.Errorf("error: utils: LoadFrom: %q: file type is missing or not supported", filename)
	}
}

func SaveTo(filename string, structure any) error {
	switch filepath.Ext(filename) {
	case ".yaml":
		return SaveToYaml(filename, structure)
	case ".json":
		return SaveToJSON(filename, structure)
	case ".ini":
		return SaveToIni(filename, structure)
	default:
		return fmt.Errorf("error: utils: SaveTo: file type is missing or not supported")
	}
}

func AppendTo(filename string, structure any) error {
	switch filepath.Ext(filename) {
	case ".yaml":
		return AppendToYaml(filename, structure)
	case ".json":
		return AppendToJSON(filename, structure)
	case ".ini":
		return AppendToIni(filename, structure)
	default:
		return fmt.Errorf("error: utils: AppendTo: file type is missing or not supported")
	}
}
