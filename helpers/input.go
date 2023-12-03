package helpers

import (
	"log"
	"os"
	"strings"
)

func ReadInput(path string) string {
	buffer, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(buffer), "\n")
}

func ReadInputToLines(path string) []string {
	content := ReadInput(path)
	lines := strings.Split(content, "\n")

	return lines
}
