package cmd

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/internal/handlers"
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"BetterPC_2.0/internal/service"
	"BetterPC_2.0/pkg/cache/localCache"
	"BetterPC_2.0/pkg/email/smtpServer"
	backendServer "BetterPC_2.0/pkg/httpServer"
	"BetterPC_2.0/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start application",
	Long:  `Initialize all dependencies, establish connections and run the server`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	logger := logging.GetLogger()
	//initializing logger

	gin.ForceConsoleColor()

	logger.Infof("Starting BetterPC 2.0 server")

	err := configs.InitConfig() //initializing config path
	if err != nil {
		logger.Fatalf("failed to initialize config: %s", err.Error())
	}

	configs.SetConfig() //setting config from yaml

	service.InitAuth(configs.GetConfig())

	warns := localCache.InitLocalCache(configs.GetConfig())
	if len(warns) > 0 {
		logger.Warn(warns)
	}

	mongoDb.MustConnectMongo(configs.GetConfig(), logger) //establishing connection to mongoDB database

	mongoDbConnection, err := mongoDb.GetMongoDB() //getting the established connection to mongoDb client and collections
	if err != nil {
		logger.Fatalf("error connecting to database: %s", err.Error())
	}

	smtpServer.InitWithConfig(configs.GetConfig())

	appRepos := repository.NewRepository(mongoDbConnection)
	appServices := service.NewService(appRepos, logger)
	appHandlers := handlers.NewHandler(appServices, logger, configs.GetConfig(), localCache.GetLocalCache())

	server := new(backendServer.Server)

	port := configs.GetConfig().Server.Port

	if err := server.Run(port, appHandlers.InitRoutes(), logger); err != nil {
		logger.Fatalf("error while running the server: %v", err.Error())
	}
}
