package middlewares

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Middleware) AdminOnly() gin.HandlerFunc {
	adminRole := m.roles.AdminRole

	return func(c *gin.Context) {
		status, err := m.CheckUserPermissions(c, adminRole)
		if err != nil {
			switch status {
			case http.StatusForbidden:
				message := fmt.Sprintf("Access denied! The user does not have permissions to proceed!")
				responseManager.WarnResponseWithLog(c, status, err.Error(), message)
			case http.StatusUnauthorized:
				message := fmt.Sprintf("Access denied! Unauthorized user!")
				responseManager.WarnResponseWithLog(c, status, err.Error(), message)
			default:
				responseManager.ErrorResponseWithLog(c, status, err.Error())
			}
			
			return
		}

		c.Next()
	}
}
