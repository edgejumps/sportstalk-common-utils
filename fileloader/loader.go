package fileloader

import (
	"encoding/json"
	"os"
)

func WriteInto(target interface{}, path string, ensureDir bool) error {
	if ensureDir {
		if err := EnsureDirExist(path); err != nil {
			return err
		}
	}

	bytes, err := json.Marshal(target)

	if err != nil {
		return err
	}

	err = os.WriteFile(path, bytes, 0644)

	if err != nil {
		return err
	}

	return nil
}
