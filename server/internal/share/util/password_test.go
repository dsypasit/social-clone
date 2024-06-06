package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGeneratePassword(t *testing.T) {
	input := "test1234"
	want, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	assert.Nilf(t, err, "Unexpected error: %v", err)

	actual, err := GeneratePassword(input)

	assert.Nilf(t, err, "Unexpected error: %v", err)
	assert.Equalf(t, len(string(want)), len(actual), "Want %v but got %v", len(string(want)), len(actual))
}

func TestVerifyPassword(t *testing.T) {
	input := "test1234"
	alreadyPass, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	assert.Nilf(t, err, "Unexpected error: %v", err)

	verify, err := VerifyPassword(input, string(alreadyPass))
	assert.Nilf(t, err, "Unexpected error: %v", err)
	assert.Equal(t, true, verify, "Expected %v but got %v", true, verify)
}
