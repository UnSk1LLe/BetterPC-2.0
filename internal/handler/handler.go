package handler

import (
	"BetterPC_2.0/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	//verify := router.POST("/:token")

	/*passwordRecovery := router.Group("/password_recovery")
	{
		passwordRecovery.POST("/:email")
		passwordRecovery.POST("/:token")
	}

	shop := router.Group("/shop")
	{
		categories := shop.Group("/categories")
		{
			categories.GET("/")
			categories.GET("/:category_name")
			categories.GET("/:category_name/:product_id")
			categories.POST("/")

			products := categories.Group("/products")
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

	users := shop.Group("/users")
	{
		users.GET("/")
		users.GET("/:user_id")
		users.POST("/")
		users.PUT("/")
		users.DELETE("/")
	}

	roles := shop.Group("/roles")
	{
		roles.GET("/")
		roles.GET("/:role_id")
		roles.POST("/")
		roles.PUT("/")
		roles.DELETE("/")
	}*/

	return router
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}
