package main

import (
	BetterPC_2_0 "BetterPC_2.0"
	"BetterPC_2.0/configs"
	"BetterPC_2.0/internal/handler"
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/internal/service"
	"BetterPC_2.0/pkg/logging"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	logger := logging.GetLogger()

	logger.Infof("Starting BetterPC 2.0 server")

	err := configs.InitConfig()
	if err != nil {
		logger.Fatalf("failed to initialize config: %s", err.Error())
	}

	configs.SetConfig()

	fmt.Println(configs.GetConfig())

	mongoConn, err := repository.Init(configs.GetConfig(), logger)
	if err != nil {
		logger.Fatalf("error connecting to database: %s", err.Error())
	}

	logrus.Infof("asd")
	repos := repository.NewRepository(mongoConn)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(BetterPC_2_0.Server)

	if err := server.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logger.Fatalf("error while running the server: %v", err.Error())
	}
}
