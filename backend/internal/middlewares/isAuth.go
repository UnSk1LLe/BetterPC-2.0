package middlewares

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func (m *Middleware) IsAuthorized(c *gin.Context) {
	user, exists := c.Get(UserCtx)
	if !exists || user == nil {
		message := "no user context found"
		m.logger.Error(message)
		responseManager.ErrorResponseWithLog(c, http.StatusUnauthorized, message)
		return
	}

	_, ok := user.(userResponses.UserResponse)
	if !ok {
		message := fmt.Sprintf("failed type assertion from <%v> to <%v>",
			reflect.TypeOf(user), reflect.TypeOf(userResponses.UserResponse{}),
		)
		m.logger.Error(message)
		responseManager.ErrorResponseWithLog(c, http.StatusUnauthorized, message)
		return
	}

	c.Next()
}
