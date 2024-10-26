package errors

import (
	"BetterPC_2.0/pkg/html"
	"github.com/gin-gonic/gin"
)

func RenderError(c *gin.Context, status int, action string, method string, err error, message ...string) {
	defaultMessage := "Server could not proceed your request."

	// Use provided message if exists, otherwise use default
	finalMessage := defaultMessage

	if len(message) > 0 && message[0] != "" {
		finalMessage = message[0]
	}

	html.Render(c, status, "templates/pages/error", gin.H{
		"Action":       action,
		"ActionMethod": method,
		"Error":        err.Error(),
		"Message":      finalMessage,
	})
}
