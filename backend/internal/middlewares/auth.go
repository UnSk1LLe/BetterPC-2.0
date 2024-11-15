package middlewares

import (
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (m *Middleware) UserIdentity() gin.HandlerFunc {

	return func(c *gin.Context) {

		var response userResponses.UserResponse

		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				m.logger.Errorf("error getting access token: %v", err)
			}
		}

		if err == nil && accessToken != "" {
			response, err = m.services.ParseAccessToken(accessToken)
			if err != nil {
				m.logger.Errorf("error parsing access token: %v", err)
			} else {
				c.Set(UserCtx, response)
				c.Next()
				return
			}
		}

		refreshToken, err := c.Cookie("refreshToken")
		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				m.logger.Errorf("error getting refresh token: %v", err)
			}

			m.removeTokens(c)
			c.Next()
			return
		}

		user, tokens, err := m.services.RefreshTokens(refreshToken)
		if err != nil {
			m.logger.Errorf("error refreshing tokens: %v", err)
			m.removeTokens(c)
			c.Next()
			return
		}

		c.SetCookie("accessToken", tokens.AccessToken, 15*60, "/", "", true, true)
		c.SetCookie("refreshToken", tokens.RefreshToken, 7*24*60*60, "/", "", true, true)
		response = user

		m.logger.Info("tokens refreshed successfully")

		c.Set(UserCtx, response)
		c.Next()
	}
}
