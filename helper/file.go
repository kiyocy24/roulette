package helper

import (
	"os"
	"path"
)

func WriteFile(filepath string, data []byte) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		err := os.MkdirAll(path.Dir(filepath), 0777)
		if err != nil {
			return err
		}
	}
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(filepath string) ([]byte, error) {
	return os.ReadFile(filepath)
}
