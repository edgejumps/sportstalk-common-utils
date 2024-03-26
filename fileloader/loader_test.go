package fileloader

import (
	"testing"
)

const (
	dir  = "./test"
	path = "./test/test.json"
)

func TestEnsureDirExist(t *testing.T) {
	err := EnsureDirExist(dir)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestWriteInto(t *testing.T) {

	data := map[string]interface{}{
		"test": "test",
		"nested": map[string]interface{}{
			"test": "test",
		},
	}

	err := WriteInto(data, path, true)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestReadAsMap(t *testing.T) {

	data, err := ReadAsMap(path)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if data == nil {
		t.Errorf("Expected not nil, got nil")
	}

	if data["test"] != "test" {
		t.Errorf("Expected test, got %v", data["test"])
	}

	if data["nested"].(map[string]interface{})["test"] != "test" {
		t.Errorf("Expected test, got %v", data["nested"].(map[string]interface{})["test"])
	}
}

func TestReadInto(t *testing.T) {

	var data map[string]interface{}

	err := ReadInto(path, &data)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if data["test"] != "test" {
		t.Errorf("Expected test, got %v", data["test"])
	}

	if data["nested"].(map[string]interface{})["test"] != "test" {
		t.Errorf("Expected test, got %v", data["nested"].(map[string]interface{})["test"])
	}
}

func TestReadPropertiesFile(t *testing.T) {

	data, err := ReadPropertiesFile("./test.properties")

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if data == nil {
		t.Errorf("Expected not nil, got nil")
	}

	if data["key"] != "value" {
		t.Errorf("Expected value, got %v", data["key"])
	}
}

func TestCheckFile(t *testing.T) {

	existed := CheckFile(path)

	if !existed {
		t.Errorf("Expected true, got false")
	}

	existed = CheckFile("./test/notexist.json")

	if existed {
		t.Errorf("Expected false, got true")
	}
}
