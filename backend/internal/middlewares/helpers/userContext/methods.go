package userContext

import (
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"reflect"
)

const UserCtx = "User"

func GetUserCtx(c *gin.Context) (userResponses.UserResponse, error) {
	var response userResponses.UserResponse

	user, ok := c.Get(UserCtx)
	if !ok {
		return response, errors.New("no user context found")
	}

	response, ok = user.(userResponses.UserResponse)
	if !ok {
		errMessage := fmt.Sprintf("error getting user: type assertion failed, from <%v> to <%v>", reflect.TypeOf(user), reflect.TypeOf(userResponses.UserResponse{}))
		return response, errors.New(errMessage)
	}

	return response, nil
}

func SetUserCtx(c *gin.Context, user userResponses.UserResponse) {
	c.Set(UserCtx, user)
}
