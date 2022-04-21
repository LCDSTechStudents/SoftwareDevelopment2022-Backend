package crypto

import "golang.org/x/crypto/bcrypt"

type BCrypt struct {
	SignKey []byte
}

func (B *BCrypt) HashPassword(password string) string {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashBytes)
}

func (B *BCrypt) CheckPW(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false
	}
	return true
}
