package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	password, limit := []byte(p), 8

	check, _ := bcrypt.GenerateFromPassword(password, limit)

	return string(check)
}

func ComparePass(h, p []byte) bool {
	check, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(check, pass)

	return err == nil
}
