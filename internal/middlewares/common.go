package middlewares

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/internal/service"
	"BetterPC_2.0/pkg/data/models/users"
	"BetterPC_2.0/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"reflect"
)

const (
	UserCtx = "User"
)

type Middleware struct {
	services *service.Service
	logger   *logging.Logger
	roles    *configs.UserRoles
}

func NewMiddleware(services *service.Service, logger *logging.Logger, cfg *configs.Config) *Middleware {
	return &Middleware{
		services: services,
		logger:   logger,
		roles:    &cfg.User.Roles,
	}
}

func (m *Middleware) removeTokens(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "", true, true)
	c.SetCookie("refreshToken", "", -1, "/", "", true, true)
}

func (m *Middleware) CheckUserPermissions(c *gin.Context, roles ...string) (status int, err error) {
	user, ok := c.Get(UserCtx)
	if !ok {
		m.logger.Warn("no user context found")
		return http.StatusBadRequest, errors.New("error getting user")
	}

	response, ok := user.(users.UserResponse)
	if !ok {
		m.logger.Errorf("type assertion failed, from <%v> to <%v>", reflect.TypeOf(user), reflect.TypeOf(users.UserResponse{}))
		return http.StatusBadRequest, errors.New("error getting user")
	}

	if response.Role != m.roles.AdminRole {
		m.logger.Warnf("access denied for non-admin user: %s", response.ID)
		return http.StatusForbidden, errors.New("Access denied! The user does not have permissions to proceed!")
	}

	userObjId, err := primitive.ObjectIDFromHex(response.ID)
	if err != nil {
		m.logger.Errorf("error converting user object id to primitive.ObjectID: %v", err)
		return http.StatusInternalServerError, errors.New(fmt.Sprintf("error converting user object id to primitive.ObjectID: %v", err))
	}

	hasRole, err := m.services.HasRole(userObjId, roles...)
	if !hasRole {
		m.logger.Warnf("access denied for non-admin user: %s", response.ID)
		return http.StatusForbidden, errors.New(fmt.Sprintf("Access denied! The user does not have permissions to proceed: %v", err))
	}

	return http.StatusOK, nil
}
