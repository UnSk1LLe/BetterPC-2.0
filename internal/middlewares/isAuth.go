package middlewares

import (
	"BetterPC_2.0/pkg/data/models/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func (m *Middleware) IsAuthorized(c *gin.Context) {
	user, exists := c.Get(UserCtx)
	if !exists || user == nil {
		m.logger.Error("no user context found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no user context found"})
		return
	}

	_, ok := user.(users.UserResponse)
	if !ok {
		m.logger.Errorf("type assertion failed, from <%v> to <%v>", reflect.TypeOf(user), reflect.TypeOf(users.UserResponse{}))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("failed type assertion from <%v> to <%v>",
			reflect.TypeOf(user), reflect.TypeOf(users.UserResponse{})),
		})
		return
	}

	c.Next()
}
