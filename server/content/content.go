package content

import (
	"SoftwareDevelopment-Backend/config"
	"go.uber.org/zap"
)

type Content struct {
	Log    *zap.Logger
	Config *config.Config
	Data   []interface{}
}

func InitContent(config *config.Config, log *zap.Logger, service int, data ...interface{}) *Content {
	return &Content{
		Config: config,
		Log:    log,
		Data:   data,
	}
}
