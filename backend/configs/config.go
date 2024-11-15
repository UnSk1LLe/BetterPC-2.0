package configs

import (
	"github.com/spf13/viper"
	"time"
)

var Configurations *Config

type Config struct {
	App           App
	Server        Server
	MongoDB       MongoDB
	LocalCache    LocalCache
	Tokens        Tokens
	Notifications Notifications
	User          User
}

type App struct {
	Name string
}

type Server struct {
	Url  string
	Port string
}

type MongoDB struct {
	Url                   string
	UsersDbName           string
	ShopDbName            string
	UsersCollectionsNames []string
	ShopCollectionsNames  []string
}

type LocalCache struct {
	ExpirationTime time.Duration
	PurgeTime      time.Duration
}

type Tokens struct {
	AccessTokenTTL         time.Duration
	AccessTokenSigningKey  string
	RefreshTokenTTL        time.Duration
	RefreshTokenSigningKey string
	VerificationTokenTTL   time.Duration
}

type Notifications struct {
	Email    string
	Password string
	SmtpHost string
	SmtpPort string
}

type User struct {
	Image string
	Roles UserRoles
}

type UserRoles struct {
	CustomerRole      string
	ShopAssistantRole string
	AdminRole         string
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
	Configurations = &Config{
		MongoDB: MongoDB{
			Url:                   viper.GetString("mongoDb.url"),
			UsersDbName:           viper.GetString("mongoDb.usersDbName"),
			UsersCollectionsNames: viper.GetStringSlice("mongoDb.usersDbCollection"),
			ShopDbName:            viper.GetString("mongoDb.shopDbName"),
			ShopCollectionsNames:  viper.GetStringSlice("mongoDb.shopDbCollections"),
		},
		Server: Server{
			Url:  viper.GetString("server.url"),
			Port: viper.GetString("server.port"),
		},
		App: App{
			Name: viper.GetString("app.name"),
		},
		LocalCache: LocalCache{
			ExpirationTime: viper.GetDuration("localCache.expirationTime"),
			PurgeTime:      viper.GetDuration("localCache.purgeTime"),
		},
		Tokens: Tokens{
			AccessTokenTTL:         viper.GetDuration("tokens.accessTokenTTL"),
			AccessTokenSigningKey:  viper.GetString("tokens.accessTokenSigningKey"),
			RefreshTokenTTL:        viper.GetDuration("tokens.refreshTokenTTL"),
			RefreshTokenSigningKey: viper.GetString("tokens.refreshTokenSigningKey"),
			VerificationTokenTTL:   viper.GetDuration("tokens.verificationTokenTTL"),
		},
		Notifications: Notifications{
			Email:    viper.GetString("notifications.email"),
			Password: viper.GetString("notifications.password"),
			SmtpHost: viper.GetString("notifications.smtpHost"),
			SmtpPort: viper.GetString("notifications.smtpPort"),
		},
		User: User{
			Image: viper.GetString("users.image"),
			Roles: UserRoles{
				CustomerRole:      viper.GetString("users.roles.customerRole"),
				ShopAssistantRole: viper.GetString("users.roles.shopAssistantRole"),
				AdminRole:         viper.GetString("users.roles.adminRole"),
			},
		},
	}
}

func GetConfig() *Config {
	if Configurations == nil {
		panic("config not set")
	}
	return Configurations
}