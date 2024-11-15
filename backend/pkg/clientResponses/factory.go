package clientResponses

import (
	"BetterPC_2.0/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

type ClientResponseWithLog func(c *gin.Context, status int, logMessage string, respMessage ...string)
type ClientResponse func(c *gin.Context, status int, respMessage ...string)

func NewResponse(defaultMessage, jsonFieldName string, ginMethod func(*gin.Context, int, any)) ClientResponse {
	return func(c *gin.Context, status int, respMessage ...string) {
		finalMessage := defaultMessage

		if len(respMessage) > 0 {
			if len(respMessage) == 1 {
				finalMessage = respMessage[0]
			} else {
				finalMessage = strings.Join(respMessage, ", ")
			}
		}

		ginMethod(c, status, gin.H{jsonFieldName: finalMessage})
	}
}

func NewResponseWithLog(jsonFieldName string, ginMethod func(*gin.Context, int, any), logLevel logrus.Level) ClientResponseWithLog {
	return func(c *gin.Context, status int, logMessage string, respMessage ...string) {
		logger := logging.GetLogger()

		var finalMessage string

		if len(respMessage) > 0 {
			if len(respMessage) == 1 {
				finalMessage = respMessage[0]
			} else {
				finalMessage = strings.Join(respMessage, ", ")
			}
		} else {
			finalMessage = logMessage
		}

		logger.Log(logLevel, logMessage)

		ginMethod(c, status, gin.H{jsonFieldName: finalMessage})
	}
}
