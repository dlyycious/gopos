package authmodule

import "golang.org/x/crypto/bcrypt"

func PasswordCheck(passwordHash []byte, pasword []byte) bool {
	if err := bcrypt.CompareHashAndPassword(passwordHash, pasword); err != nil {
		return false
	}
	return true
}
