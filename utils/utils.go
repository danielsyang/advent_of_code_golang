package utils

import (
	"bufio"
	"fmt"
	"os"

	"path/filepath"
)

func ReadFile(fileName string) []string {

	dir, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintln("Could not get current directory", fileName))
	}

	path := filepath.Join(dir, fileName)

	file, err := os.Open(path)

	if err != nil {
		panic(fmt.Sprintln("Could not open file", path))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := []string{}

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintln("Error scanning file", err))
	}

	return result

}

func AbsInt(number int) int {

	if number >= 0 {
		return number
	}

	return -number

}
