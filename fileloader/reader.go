package fileloader

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	ErrFailedDecodeAsMap = fmt.Errorf("failed to decode json file as map")
)

func ReadAsMap(path string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	data := make(map[string]interface{})

	err = json.NewDecoder(jsonFile).Decode(&data)

	if err != nil {
		return nil, ErrFailedDecodeAsMap
	}

	return data, nil
}

func ReadInto(path string, target interface{}) error {
	jsonFile, err := os.Open(path)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	return json.NewDecoder(jsonFile).Decode(target)
}
