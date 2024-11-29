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
	Stripe        Stripe
}

type App struct {
	Name string
}

type Server struct {
	Url  string
	Port string
}

type MongoDB struct {
	Username                string
	Password                string
	ClusterAddress          string
	Options                 string
	UsersDbName             string
	ShopDbName              string
	UsersCollectionName     string
	OrdersCollectionName    string
	ProductsCollectionNames []string
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
	Roles UserRoles
}

type UserRoles struct {
	CustomerRole      string
	ShopAssistantRole string
	AdminRole         string
}

type Stripe struct {
	PublicKey  string
	PrivateKey string
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
			Username:                viper.GetString("mongoDb.username"),
			Password:                viper.GetString("mongoDb.password"),
			ClusterAddress:          viper.GetString("mongoDb.clusterAddress"),
			Options:                 viper.GetString("mongoDb.options"),
			UsersDbName:             viper.GetString("mongoDb.usersDbName"),
			UsersCollectionName:     viper.GetString("mongoDb.usersCollection"),
			ShopDbName:              viper.GetString("mongoDb.shopDbName"),
			OrdersCollectionName:    viper.GetString("mongoDb.ordersCollection"),
			ProductsCollectionNames: viper.GetStringSlice("mongoDb.productsCollectionList"),
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
			Roles: UserRoles{
				CustomerRole:      viper.GetString("users.roles.customerRole"),
				ShopAssistantRole: viper.GetString("users.roles.shopAssistantRole"),
				AdminRole:         viper.GetString("users.roles.adminRole"),
			},
		},
		Stripe: Stripe{
			PublicKey:  viper.GetString("stripe.publicKey"),
			PrivateKey: viper.GetString("stripe.privateKey"),
		},
	}
}

func GetConfig() *Config {
	if Configurations == nil {
		panic("config not set")
	}
	return Configurations
}
