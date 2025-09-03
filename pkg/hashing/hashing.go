package hashing

import "golang.org/x/crypto/bcrypt"

func Hash(str string) string {
	pass := []byte(str)
	hashedStrByte, err := bcrypt.GenerateFromPassword(pass, 0)
	var hashedStr string
	if err != nil {
		hashedStr = string(hashedStrByte)
	}

	return hashedStr
}
