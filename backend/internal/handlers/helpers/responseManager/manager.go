package responseManager

import (
	"BetterPC_2.0/pkg/clientResponses"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	DefaultError         = "Server could not proceed your request"
	ErrorResponseWithLog = clientResponses.NewResponseWithLog(
		"error",
		func(c *gin.Context, code int, jsonObj any) {
			c.AbortWithStatusJSON(code, jsonObj)
		},
		logrus.ErrorLevel,
	)
	ErrorResponse = clientResponses.NewResponse(
		DefaultError,
		"error",
		func(c *gin.Context, code int, jsonObj any) {
			c.AbortWithStatusJSON(code, jsonObj)
		},
	)
	DefaultWarn         = "Something went wrong, try again later"
	WarnResponseWithLog = clientResponses.NewResponseWithLog(
		"warn",
		func(c *gin.Context, code int, jsonObj any) {
			c.AbortWithStatusJSON(code, jsonObj)
		},
		logrus.WarnLevel,
	)
	DefaultMessage         = "Request completed successfully"
	MessageResponseWithLog = clientResponses.NewResponseWithLog(
		"message",
		func(c *gin.Context, code int, jsnObj any) {
			c.JSON(code, jsnObj)
		},
		logrus.InfoLevel,
	)
	MessageResponse = clientResponses.NewResponse(
		DefaultMessage,
		"message",
		func(c *gin.Context, code int, jsnObj any) {
			c.JSON(code, jsnObj)
		},
	)
)
