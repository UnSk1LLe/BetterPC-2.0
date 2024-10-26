package html

import (
	"github.com/gin-gonic/gin"
)

func LoadTemplates(router *gin.Engine) {
	router.LoadHTMLGlob("internal/templates/**/*tmpl")
}
