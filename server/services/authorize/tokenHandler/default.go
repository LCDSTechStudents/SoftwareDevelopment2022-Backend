package tokenHandler

import (
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
)

type Claims struct {
	UserId uint64
	jwt.StandardClaims
}

type DefaultJWT struct {
	log     *zap.Logger
	signKey string
}

func (d *DefaultJWT) GenerateToken(id uint64) string {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(d.signKey)
	if err != nil {
		d.log.Error("signing key error", zap.Error(err))
		return ""
	}
	return tokenString
}

func (d *DefaultJWT) VerifyToken(s string) int {
	token, err := d.parseToken(s)
	if err != nil {
		return Error
	}
	if token.Valid {
		return ValidToken
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return NotAToken
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return ExpiredToken
		} else {
			return CouldNotHandle
		}
	} else {
		return CouldNotHandle
	}
}

func (d *DefaultJWT) parseToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(d.signKey), nil
	})
	return token, err

}
