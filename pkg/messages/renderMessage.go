package messages

import (
	"BetterPC_2.0/pkg/html"
	"github.com/gin-gonic/gin"
)

func RenderMessage(c *gin.Context, status int, action string, method string, message string) {
	html.Render(c, status, "templates/messages/message", gin.H{
		"Action":       action,
		"ActionMethod": method,
		"Message":      message,
	})
}
