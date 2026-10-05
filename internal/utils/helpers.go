package utils

import (
	"fmt"
	"path/filepath"
)

func LoadFrom(filename string, structure interface{}) error {
	switch filepath.Ext(filename) {
	case ".yaml":
		return LoadFromYaml(filename, structure)
	case ".json":
		return LoadFromJSON(filename, structure)
	case ".ini":
		return LoadFromIni(filename, structure)
	default:
		return fmt.Errorf("error: utils: LoadFrom: file type is missing or not supported")
	}
}

func SaveTo(filename string, structure interface{}) error {
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
