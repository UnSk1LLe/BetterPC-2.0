package middlewares

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/internal/service"
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"BetterPC_2.0/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"slices"
)

const (
	UserCtx = "User"
)

type Middleware struct {
	services *service.Service
	logger   *logging.Logger
	roles    *configs.UserRoles
	cache    *cache.Cache
}

func NewMiddleware(services *service.Service, logger *logging.Logger, cfg *configs.Config, lCache *cache.Cache) *Middleware {
	return &Middleware{
		services: services,
		logger:   logger,
		roles:    &cfg.User.Roles,
		cache:    lCache,
	}
}

func (m *Middleware) removeTokens(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "", true, true)
	c.SetCookie("refreshToken", "", -1, "/", "", true, true)
}

// CheckUserPermissions can only be used in chain after UserIdentity middleware
func (m *Middleware) CheckUserPermissions(c *gin.Context, roles ...string) (status int, err error) {
	user, ok := c.Get(UserCtx)
	if !ok {
		return http.StatusUnauthorized, errors.New("no user context found")
	}

	response, ok := user.(userResponses.UserResponse)
	if !ok {
		errMessage := fmt.Sprintf("error getting user: type assertion failed, from <%v> to <%v>", reflect.TypeOf(user), reflect.TypeOf(userResponses.UserResponse{}))
		return http.StatusBadRequest, errors.New(errMessage)
	}

	if !slices.Contains(roles, response.Role) {
		errMessage := fmt.Sprintf("access denied for non-admin user: %s", response.ID)
		return http.StatusForbidden, errors.New(errMessage)
	}

	hasRole, err := m.services.HasRole(response.ID, roles...)
	if err != nil {
		errMessage := fmt.Sprintf("access denied! error confirming user permissions: %v", err)
		return http.StatusInternalServerError, errors.New(errMessage)
	}
	if !hasRole {
		errMessage := fmt.Sprintf("access denied for non-admin user: %s", response.ID)
		return http.StatusForbidden, errors.New(errMessage)
	}

	return http.StatusOK, nil
}
