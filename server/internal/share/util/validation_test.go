package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTEmailValidation(t *testing.T) {
	testTable := []struct {
		title string
		input string
		want  bool
	}{
		{"should valid", "a@gmail.com", true},
		{"should invalid", "a.gmail.com", false},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			actual := ValidateEmail(v.input)
			assert.Equalf(t, v.want, actual, "want %v but got %v", v.want, actual)
		})
	}
}
