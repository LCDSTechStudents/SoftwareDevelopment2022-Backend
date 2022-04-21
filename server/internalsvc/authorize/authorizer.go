package authorize

import (
	"SoftwareDevelopment-Backend/config"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/crypto"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/idGenerator"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/smtp"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/tokenHandler"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/userpack"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/verifyCodeHandler"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	AUTHORIZER    = "authorizer"
	NAME          = "AUTH"
	OK            = 1
	WrongPassword = 2
	UserNotFound  = 3
)

type DefaultAuthorizer struct {
	*zap.Logger
	config *config.Config
	db     *gorm.DB
	idGenerator.IDGenerator
	tokenHandler.TokenHandler
	smtp.EmailHandler
	crypto.PasswordHandler
	verifyCodeHandler.VerifyCodeHandler
}

func (d *DefaultAuthorizer) Run() error {
	return nil
}

func (d *DefaultAuthorizer) GetDB() *gorm.DB {
	return d.db
}

func InitDefaultAuthorizer(log *zap.Logger, config *config.Config) IAuthorizer {
	dsn := getDSN(config.Services.Auth.DB, log)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}
	r := &DefaultAuthorizer{
		Logger:            log,
		config:            config,
		db:                db,
		IDGenerator:       idGenerator.InitDefaultIDGenerator(),
		TokenHandler:      tokenHandler.InitTokenHandler(log, config),
		EmailHandler:      smtp.InitDefaultSMTP(log, config),
		PasswordHandler:   crypto.InitPasswordHandler(config),
		VerifyCodeHandler: verifyCodeHandler.InitDefaultCodeHandler(log, config),
	}
	return r
}

func (d *DefaultAuthorizer) VerifyLoginInfo(email string, pw string) (*userpack.User, int) {
	var user *userpack.User
	d.db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, UserNotFound
	}
	if !d.CheckPW(pw, user.Password) {
		return nil, WrongPassword
	}
	return user, OK
}

func getDSN(DB config.DB, log *zap.Logger) string {
	un := DB.UserName
	pc := DB.Password
	prtc := DB.Protocol
	url := DB.URL
	dn := DB.DBName
	r := un + ":" + pc + "@" + prtc + "(" + url + ")/" + dn
	log.Info("getting dsn: " + r)
	return r
}
