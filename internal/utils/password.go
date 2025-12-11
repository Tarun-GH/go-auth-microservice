package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//bcrypt.DefaultCost is 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
