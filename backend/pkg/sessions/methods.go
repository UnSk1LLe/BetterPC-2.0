package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context, key, value string) {
	session := sessions.Default(c)

	session.Set(key, value)
	session.Save()
}

func Flash(c *gin.Context, key string) string {
	session := sessions.Default(c)

	response := session.Flashes(key)
	session.Save()
	return response[0].(string)
}

func Get(c *gin.Context, key string) string {
	session := sessions.Default(c)

	response := session.Get(key)
	session.Save()
	return response.(string)
}

func Remove(c *gin.Context, key string) {
	session := sessions.Default(c)

	session.Delete(key)
	session.Save()
}
