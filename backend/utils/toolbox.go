package utils

import "io/ioutil"

// WriteFile 将json写入文件
func WriteFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	return err
}

func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return data, err
}
