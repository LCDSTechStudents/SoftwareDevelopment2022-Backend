package crypto

import "SoftwareDevelopment-Backend/config"

type PasswordHandler interface {
	HashPassword(password string) string
	Check(hashed string, hash string) bool
}

func InitPasswordHandler(config *config.Config) PasswordHandler {
	return &BCrypt{
		SignKey: []byte(config.Services.Auth.PasswordCrypt),
	}
}
