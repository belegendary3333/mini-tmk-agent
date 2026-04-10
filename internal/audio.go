package internal

import (
	"os"
)

func SaveResult(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
