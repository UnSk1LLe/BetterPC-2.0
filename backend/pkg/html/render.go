package html

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Render(c *gin.Context, status int, name string, data gin.H) {
	c.HTML(status, name, withGlobalData(c, data))
}

func withGlobalData(c *gin.Context, data gin.H) gin.H {
	data["appName"] = viper.Get("app.name")
	//data["User"] = getUserData(c)
	return data
}

/*func getUserData(c *gin.Context) userResponses.UserResponse {
	user, ok := c.Get(middlewares.UserCtx)
	if !ok {
		return userResponses.UserResponse{}
	}

	response, ok := user.(userResponses.UserResponse)
	if !ok {
		return userResponses.UserResponse{}
	}

	return response
}*/
