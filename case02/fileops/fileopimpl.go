package fileops

import (
	"os"
)

func WriteFileOp(fileName string, content string) {
	os.WriteFile(fileName, []byte(content), 0644)
}

func ReadFileOp(fileName string) (string, error) {
	content, error := os.ReadFile(fileName)
	var stringContent = string(content)
	if error != nil {
		return "", error
	}
	return stringContent, nil
}
