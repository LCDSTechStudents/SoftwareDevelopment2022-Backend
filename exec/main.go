package main

import (
	"SoftwareDevelopment-Backend/config"
	"SoftwareDevelopment-Backend/logger"
	"SoftwareDevelopment-Backend/server"
	_default "SoftwareDevelopment-Backend/server/internalsvc/recognition/default"
	"fmt"
)

func main() {
	//s := initDependencies()
	//s.Run()
	fmt.Println(_default.ToBinary("./files/refacto/space_shuttle.jpg"))

}

func initDependencies() server.Server {
	log := logger.InitLogger()
	config := config.InitConfig(log)
	return server.InitHTTPServer(config, log)

}
