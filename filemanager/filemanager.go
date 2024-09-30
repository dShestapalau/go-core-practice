package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("Could not open file %v\n", path)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Failed to read content from file.")
		return nil, err
	}

	file.Close()
	return lines, nil
}
