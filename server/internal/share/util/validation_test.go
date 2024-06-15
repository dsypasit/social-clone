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
			actual := IsValidEmail(v.input)
			assert.Equalf(t, v.want, actual, "want %v but got %v", v.want, actual)
		})
	}
}

func TestValidateUUID(t *testing.T) {
	testTable := []struct {
		title string
		input string
		want  bool
	}{
		{"should valid", "8cef4343-1484-4f96-be14-87f235fb492", false},
		{"should invalid", "8cef4343-1484-4f96-be14-87f235fb4924", true},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			actual := IsValidUUID(v.input)
			assert.Equalf(t, v.want, actual, "want %v but got %v", v.want, actual)
		})
	}
}
