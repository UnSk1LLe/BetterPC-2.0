package middlewares

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	DefaultRateLimit    = 5
	DefaultRateInterval = 2 * time.Minute
)

func (m *Middleware) RateLimitFromClient(limit int, interval time.Duration) gin.HandlerFunc {

	if limit <= 0 {
		m.logger.Warnf("limit must be greater than zero; using default limit: %d", DefaultRateLimit)
		limit = DefaultRateLimit
	}
	if interval == 0 || interval > 10*time.Minute {
		m.logger.Warnf("interval must be in between 0s and 10m0s; using default interval: %s", DefaultRateInterval)
		interval = DefaultRateInterval
	}

	return func(c *gin.Context) {
		cacheKey := fmt.Sprintf("%s_%s", c.ClientIP(), c.Request.URL.Path)

		s, ok := m.cache.Get(cacheKey)
		if !ok {
			m.cache.Set(cacheKey, 1, interval)
		} else {
			err := m.cache.Increment(cacheKey, 1)
			if err != nil {
				m.cache.Set(cacheKey, 1, interval)
				m.logger.Errorf("%s, cache value renewed", err.Error())
			}

			if s.(int)+1 > limit {
				message := "too many requests, try again later"
				logMessage := fmt.Sprintf("client with ip '%s' exceeds limit %d for '%s'", c.ClientIP(), limit, c.Request.URL.Path)
				responseManager.WarnResponseWithLog(c, http.StatusTooManyRequests, logMessage, message)
				return
			}
		}

		c.Next()
	}
}
