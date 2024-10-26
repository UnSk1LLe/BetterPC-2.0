package html

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Render(c *gin.Context, status int, name string, data gin.H) {
	c.HTML(status, name, withGlobalData(data))
}

func withGlobalData(data gin.H) gin.H {
	data["appName"] = viper.Get("app.name")

	return data
}
