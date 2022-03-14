package authorize

import (
	"SoftwareDevelopment-Backend/config"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/crypto"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/idGenerator"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/smtp"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/tokenHandler"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	NAME = "AUTH"
)

type DefaultAuthorizer struct{
	*zap.Logger
	config *config.Config
	db *gorm.DB
	idGen idGenerator.IDGenerator
	tokenHandler tokenHandler.TokenHandler
	smtp smtp.EmailHandler
	crypto crypto.PasswordHandler
}

func InitDefaultAuthorizer(log *zap.Logger, config *config.Config) IAuthorizer{
	dsn := getDSN(config.Services.Auth.DB, log)
	gorm.Open(mysql.)
	r := &DefaultAuthorizer{
		Logger:       log,
		config:       config,
		db:           db,
		idGen:        ,
		tokenHandler: nil,
		smtp:         nil,
		crypto:       nil,
	}
}

func getDSN(DB config.DB, log *zap.Logger) string {
	un := DB.UserName
	pc := DB.Password
	prtc := DB.Protocol
	url := DB.URL
	dn := DB.DBName
	r := un + ":" + pc + "@" + prtc + "(" + url + ")/" + dn
	log.Info("getting dsn: "+ r)
	return r
}

