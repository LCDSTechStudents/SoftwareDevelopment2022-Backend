package tokenHandler

import (
	"SoftwareDevelopment-Backend/config"
	"go.uber.org/zap"
)

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

func InitTokenHandler(log *zap.Logger, config *config.Config) TokenHandler {
	return &DefaultJWT{
		signKey: []byte(config.Services.Auth.JWTKey),
		log:     log,
	}
}
