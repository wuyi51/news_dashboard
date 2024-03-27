package utils

import (
	"os"
)

// WriteFile 将json写入文件
func WriteFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	return err
}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	return data, err
}
