package html

import (
	"BetterPC_2.0/internal/middlewares"
	"BetterPC_2.0/pkg/data/models/users"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Render(c *gin.Context, status int, name string, data gin.H) {
	c.HTML(status, name, withGlobalData(c, data))
}

func withGlobalData(c *gin.Context, data gin.H) gin.H {
	data["appName"] = viper.Get("app.name")
	data["User"] = getUserData(c)
	return data
}

func getUserData(c *gin.Context) users.UserResponse {
	user, ok := c.Get(middlewares.UserCtx)
	if !ok {
		return users.UserResponse{}
	}

	response, ok := user.(users.UserResponse)
	if !ok {
		return users.UserResponse{}
	}

	return response
}
