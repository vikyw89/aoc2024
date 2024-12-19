package internals

import (
	"fmt"
	"os"
)

func LoadFile(path string) string {
	fileByte, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	fileString := string(fileByte)

	return fileString
}
