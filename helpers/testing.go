package helpers

import (
	"bytes"
	"fmt"
	"testing"
)

const fstring = "\033[%vm"
const (
	BOLD       = 1
	FG_RED     = 31
	FG_GREEN   = 32
	FG_DEFAULT = 39
	BG_RED     = 41
	BG_GREEN   = 42
	BG_DEFAULT = 49
)

func Format(text string, attributes []int) string {
	var buffer bytes.Buffer
	for _, a := range attributes {
		buffer.WriteString(fmt.Sprint(a))
		buffer.WriteString(";")
	}
	return fmt.Sprintf(fstring, buffer.String())
}

func Test(t *testing.T) {
	result := Format("A string", []int{FG_RED, BG_DEFAULT, BOLD})
	expected := "\033[31;49;1mA string"
	if result != expected {
		t.Error(fmt.Sprintf("Format(\"A string\", []int{FG_RED, BG_DEFAULT, BOLD}) = %v; want %v", result, expected))
	}
}
