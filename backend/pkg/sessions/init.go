package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))
}
