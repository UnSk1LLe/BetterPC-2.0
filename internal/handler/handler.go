package handler

import (
	"BetterPC_2.0/internal/service"
	"BetterPC_2.0/pkg/html"
	"BetterPC_2.0/pkg/logging"
	"BetterPC_2.0/pkg/static"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	logger   *logging.Logger
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	html.LoadTemplates(router)
	static.LoadStatic(router)

	auth := router.Group("/auth")
	{
		auth.GET("/register", h.RegisterForm)
		auth.POST("/register", h.Register)
		auth.GET("/login", h.LoginForm)
		auth.POST("/login", h.Login)
		//auth.GET("/refresh", h.Refresh)
	}

	//verify := router.POST("/:token")

	/*passwordRecovery := router.Group("/password_recovery")
	{
		passwordRecovery.POST("/:email")
		passwordRecovery.POST("/:token")
	}*/

	/*shop := router.Group("/shop")
	{
		categories := shop.Group("/categories")
		{
			categories.GET("/")
			categories.GET("/:category_name")
			categories.GET("/:category_name/:product_id")
			categories.POST("/")

			products := categories.Group("/productTypes")
			{
				products.GET("/")
				products.GET("/:product_id")
				products.POST("/")
				products.PUT("/:product_id")
				products.DELETE("/:product_id")
			}
		}

		orders := shop.Group("/orders")
		{
			orders.GET("/")
			orders.GET("/:order_id")
			orders.POST("/")
			orders.PUT("/:order_id")
			orders.DELETE("/:order_id")
		}
	}

	users := router.Group("/users")
	{
		users.GET("/")
		users.GET("/:user_id")
		users.POST("/")
		users.PUT("/")
		users.DELETE("/")
	}

	roles := router.Group("/roles")
	{
		roles.GET("/")
		roles.GET("/:role_id")
		roles.POST("/")
		roles.PUT("/")
		roles.DELETE("/")
	}*/

	return router
}

func NewHandler(services *service.Service, logger *logging.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}
