package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFile, outputFile string) FileManager {
	return FileManager{
		InputFilePath:  inputFile,
		OutputFilePath: outputFile,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Printf("Could not open file %v\n", fm.InputFilePath)
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

func (fm FileManager) WriteJSON(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		fmt.Println("Failed to create file.")
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		fmt.Println("Failed to convert data to JSON.")
		file.Close()
		return err
	}

	file.Close()

	return nil
}
