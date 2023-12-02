package helpers

import (
	"fmt"
	"time"
)

func printResult(name string, result string, duration time.Duration) {
	fmt.Printf("part %v: %v\n   completed in %v\n\n", name, result, duration)
}

func RunPart(part func() string, name string) {
	start := time.Now()
	result := part()
	printResult(name, result, time.Since(start))
}
