package middlewares

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/internal/middlewares/helpers/userContext"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Middleware) VerifiedOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := userContext.GetUserCtx(c)
		if err != nil {
			responseManager.ErrorResponseWithLog(c, http.StatusUnauthorized, err.Error())
			return
		}
		if user.ID == "" {
			responseManager.ErrorResponseWithLog(c, http.StatusUnauthorized, "invalid user data")
			return
		}

		c.Next()
	}
}
