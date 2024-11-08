package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Middleware) StaffOnly() gin.HandlerFunc {
	adminRole := m.roles.AdminRole
	shopAssistantRole := m.roles.ShopAssistantRole

	return func(c *gin.Context) {
		status, err := m.CheckUserPermissions(c, adminRole, shopAssistantRole)
		if err != nil && status != http.StatusOK {
			c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
			return
		}

		c.Next()
	}
}
