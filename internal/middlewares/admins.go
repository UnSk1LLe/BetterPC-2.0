package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Middleware) AdminOnly() gin.HandlerFunc {
	adminRole := m.roles.AdminRole

	return func(c *gin.Context) {
		status, err := m.CheckUserPermissions(c, adminRole)
		if err != nil && status != http.StatusOK {
			c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
			return
		}

		c.Next()
	}
}
