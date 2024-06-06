package util

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(pass string) (string, error) {
	newPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(newPass), err
}

func VerifyPassword(pass string, alreadyPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(alreadyPass), []byte(pass))
	return err == nil, err
}
