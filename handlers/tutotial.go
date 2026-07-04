package handlers

import (
	"errors"
	"fmt"
	"os"
)

func FileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err != nil {
		// os.isnot exist
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func UploadFunction() {
	path := "~/Desktop/ha.txt"
	fileExists, _ := FileExists(path)
	fmt.Println("file exists", fileExists)
}
