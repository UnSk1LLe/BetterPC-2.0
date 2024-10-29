package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	accessToken, err := c.Cookie("access-token")
	if err != nil {
		h.logger.Errorf("error getting access token: %v", err)
		return
	}

	userId, err := h.services.Authorization.ParseAccessToken(accessToken)
	if errors.Is(err, jwt.ErrTokenExpired) {
		h.logger.Infof("access token expired, trying to refresh...")
		refreshToken, err := c.Cookie("refresh-token")
		if err != nil {
			h.logger.Errorf("error getting refresh token: %v", err)
			return
		}

		tokens, err := h.services.Authorization.RefreshTokens(refreshToken)
		if err != nil {
			h.logger.Errorf("error refreshing tokens: %v", err)
			return
		}

		c.SetCookie("access-token", tokens.AccessToken, 3600, "/", "", true, true)

		c.SetCookie("refresh-token", tokens.RefreshToken, 7*24*3600, "/", "", true, true)
	}

}
