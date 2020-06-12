package util

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil

}

func Compare(hash string, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
