package file

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadFile(path string) (result []string, err error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return result, err
	}

	file, err := os.Open(absolutePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
