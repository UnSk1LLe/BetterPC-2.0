package localCache

import (
	"BetterPC_2.0/configs"
	"fmt"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

var localCache *cache.Cache

const (
	defaultExpirationTime = 10 * time.Minute
	defaultPurgeTime      = 15 * time.Minute
)

func InitLocalCache(cfg *configs.Config) string {
	expirationTime := cfg.LocalCache.ExpirationTime
	purgeTime := cfg.LocalCache.PurgeTime

	var warns []string
	if expirationTime <= 0 {
		expirationTime = defaultExpirationTime
		warn := fmt.Sprintf("invalid local cache expiration time, using default %s", defaultExpirationTime.String())
		warns = append(warns, warn)
	}

	if purgeTime <= 0 {
		purgeTime = defaultPurgeTime
		warn := fmt.Sprintf("invalid local cache purge time, using default %s", defaultPurgeTime.String())
		warns = append(warns, warn)
	}

	localCache = cache.New(expirationTime, purgeTime)
	return strings.Join(warns, "\n")
}

func GetLocalCache() *cache.Cache {
	if localCache == nil {
		panic("local cache not initialized")
	}
	return localCache
}
