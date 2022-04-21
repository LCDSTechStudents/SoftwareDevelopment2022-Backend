package main

import (
	"SoftwareDevelopment-Backend/config"
	"SoftwareDevelopment-Backend/logger"
	"SoftwareDevelopment-Backend/server"
)

func main() {
	s := initDependencies()
	s.Run()

}

func initDependencies() server.Server {
	log := logger.InitLogger()
	config := config.InitConfig(log)
	return server.InitHTTPServer(config, log)

}
