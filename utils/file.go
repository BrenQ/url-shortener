package utils

import (
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
