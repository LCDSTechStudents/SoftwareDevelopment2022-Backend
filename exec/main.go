package main

import (
	"SoftwareDevelopment-Backend/config"
	"SoftwareDevelopment-Backend/server"
	"go.uber.org/zap"
)

func main() {
	s := initDependencies()
	s.Run()
}

func initDependencies() server.Server {
	log, _ := zap.NewDevelopment()
	config := config.InitConfig(log)
	return server.InitHTTPServer(config, log)
}
