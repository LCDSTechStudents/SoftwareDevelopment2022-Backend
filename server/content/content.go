package content

import (
	"SoftwareDevelopment-Backend/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Content struct {
	Log    *zap.Logger
	Config *config.Config
	Db     *gorm.DB
}

func InitContent(config *config.Config, log *zap.Logger, service int) *Content {
	dsn := getDSN(config, service, log)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("connecting to database: ", zap.Error(err))
	}
	return &Content{
		Config: config,
		Db:     db,
	}
}

func getDSN(config *config.Config, service int, log *zap.Logger) string {
	DB := config.GetServiceDB(service)
	un := DB.UserName
	pc := DB.Password
	prtc := DB.Protocol
	url := DB.URL
	dn := DB.DBName
	r := un + ":" + pc + "@" + prtc + "(" + url + ")/" + dn
	log.Info("service: " + strconv.Itoa(service) + " service dsn: " + r)
	return r
}
