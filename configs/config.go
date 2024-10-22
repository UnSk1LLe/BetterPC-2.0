package configs

import "github.com/spf13/viper"

var Configurations Config

type Config struct {
	App     App
	Server  Server
	MongoDB MongoDB
}

type App struct {
	Name string
}

type Server struct {
	Port string
}

type MongoDB struct {
	Url                   string
	UsersDbName           string
	ShopDbName            string
	UsersCollectionsNames []string
	ShopCollectionsNames  []string
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func SetConfig() {
	Configurations = Config{
		MongoDB: MongoDB{
			Url:                   viper.GetString("mongoDb.url"),
			UsersDbName:           viper.GetString("mongoDb.usersDbName"),
			UsersCollectionsNames: viper.GetStringSlice("mongoDb.usersDbCollection"),
			ShopDbName:            viper.GetString("mongoDb.shopDbName"),
			ShopCollectionsNames:  viper.GetStringSlice("mongoDb.shopDbCollections"),
		},
		Server: Server{
			Port: viper.GetString("server.port"),
		},
		App: App{
			Name: viper.GetString("app.name"),
		},
	}
}

func GetConfig() *Config {
	return &Configurations
}
