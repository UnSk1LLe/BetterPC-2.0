package configs

import (
	"github.com/spf13/viper"
	"time"
)

var Configurations Config

type Config struct {
	App     App
	Server  Server
	MongoDB MongoDB
	Tokens  Tokens
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

type Tokens struct {
	AccessTokenTTL         time.Duration
	AccessTokenSigningKey  string
	RefreshTokenTTL        time.Duration
	RefreshTokenSigningKey string
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
		Tokens: Tokens{
			AccessTokenTTL:         time.Duration(viper.GetInt("tokens.accessTokenTTL")) * time.Minute,
			AccessTokenSigningKey:  viper.GetString("tokens.accessTokenSigningKey"),
			RefreshTokenTTL:        time.Duration(viper.GetInt("tokens.refreshTokenTTL")) * time.Hour,
			RefreshTokenSigningKey: viper.GetString("tokens.refreshTokenSigningKey"),
		},
	}
}

func GetConfig() *Config {
	return &Configurations
}
