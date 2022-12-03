package day

import (
	"os"
	"strings"
)

func ReadFile(filepath string) (string, error) {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func ReadLines(filepath string) ([]string, error) {
	dat, err := ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return strings.Split(dat, "\n"), nil
}
