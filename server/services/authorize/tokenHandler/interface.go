package tokenHandler

import "SoftwareDevelopment-Backend/config"

const (
	Error          = 0
	ValidToken     = 1
	ExpiredToken   = 2
	NotAToken      = 3
	CouldNotHandle = 4
)

type TokenHandler interface {
	GenerateToken(id uint64) string
	VerifyToken(token string) int
}

func InitTokenHandler(config *config.Config) TokenHandler {
	return &DefaultJWT{
		signKey: config.Services.Auth.JWTKey,
	}
}
