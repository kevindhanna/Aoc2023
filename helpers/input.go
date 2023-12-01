package helpers

import (
	"log"
	"os"
	"strings"
)

func ReadInputToLines(path string) []string {
	buffer, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	content := strings.TrimSuffix(string(buffer), "\n")
	lines := strings.Split(content, "\n")

	return lines
}
