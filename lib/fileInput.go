package lib

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/// ReadStringFile reads a file and parses it as an array of strings
func ReadStringFile(relativePath string) ([]string, error) {
	pwd, _ := os.Getwd()
	data, err := os.ReadFile(filepath.Join(pwd, relativePath))
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

/// ReadIntegerFile reads a file containing ints - one per line
func ReadIntegerFile(relativePath string) ([]int, error) {
	arr, err := ReadStringFile(relativePath)
	if err != nil {
		return nil, err
	}

	nums := make([]int, 0)
	for _, val := range arr {
		if val != "" {
			number, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("Could not parse string as int: " + val)
			}
			nums = append(nums, number)
		}
	}

	return nums, nil
}

func ReadDelimitedFile(relativePath string, delimiter string) ([][]string, error) {
	arr, err := ReadStringFile(relativePath)
	if err != nil {
		return nil, err
	}

	lines := make([][]string, len(arr))
	for i, row := range arr {
		lines[i] = strings.Split(row, delimiter)
	}

	return lines, nil
}
