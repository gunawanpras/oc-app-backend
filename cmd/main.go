package main

import (
	"log"

	"github.com/gunawanpras/oc-app-backend/config"
	"github.com/gunawanpras/oc-app-backend/delivery/server"
	"github.com/gunawanpras/oc-app-backend/internal/setup"
)

func main() {
	// init config
	config.Init(config.WithConfigFile("config"), config.WithConfigType("yaml"))
	conf := config.Get()

	// init external services
	externalService := setup.InitExternalServices(conf)
	defer func() {
		err := externalService.Oracle.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	// init core services
	coreService := setup.InitCoreServices(conf, externalService)

	// init server
	server.Up(coreService.Handler, conf.Server)
}
