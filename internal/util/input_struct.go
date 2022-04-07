package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type InputFileFormat int64

const (
	YAML InputFileFormat = iota
	JSON
)

type UnsupportedInputFileFormat struct {
	FileExtension string
}

func (e UnsupportedInputFileFormat) Error() string {
	return fmt.Sprintf("Unsupported input file format for file extension: %s. File extension must be either '.yaml' or '.json'.", e.FileExtension)
}

func ReadInputStructFromFile(file string, outStruct interface{}) error {
	format, err := DetectInputFileFormat(file)
	if err != nil {
		return err
	}

	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	log.Printf("Loaded input struct file: %s", file)

	switch format {
	case JSON:
		return json.Unmarshal(fileBytes, outStruct)
	case YAML:
		// convert YAML to JSON
		// this is a work-around to make the `json` tags on the struct work
		// unfortunately the go-yaml does not support `json` struct tags
		// see: https://github.com/go-yaml/yaml/issues/424
		yamlMap := make(map[string]interface{})
		err := yaml.Unmarshal(fileBytes, yamlMap)
		if err != nil {
			return err
		}
		jsonBytes, err := json.Marshal(yamlMap)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonBytes, outStruct)
	}
	return nil
}

func DetectInputFileFormat(file string) (InputFileFormat, error) {
	ext := filepath.Ext(file)
	switch ext {
	case ".yaml":
		return YAML, nil
	case ".json":
		return JSON, nil
	default:
		return -1, UnsupportedInputFileFormat{FileExtension: ext}
	}
}
