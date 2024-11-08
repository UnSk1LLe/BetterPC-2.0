package handlers

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/internal/middlewares"
	"BetterPC_2.0/internal/service"
	"BetterPC_2.0/pkg/html"
	"BetterPC_2.0/pkg/logging"
	"BetterPC_2.0/pkg/sessions"
	"BetterPC_2.0/pkg/static"
	"BetterPC_2.0/pkg/templateFunctions"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	logger   *logging.Logger
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(gin.Logger())

	sessions.Init(router)

	router.SetFuncMap(templateFunctions.TmplFuncs)
	html.LoadTemplates(router)
	static.LoadStatic(router)

	middleware := middlewares.NewMiddleware(h.services, h.logger, configs.GetConfig())

	auth := router.Group("/auth")
	{
		auth.GET("/register", h.RegisterForm)
		auth.POST("/register", h.Register)
		auth.GET("/login", h.LoginForm)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
	}

	//verify := router.POST("/:token")

	/*passwordRecovery := router.Group("/password_recovery")
	{
		passwordRecovery.POST("/:email")
		passwordRecovery.POST("/:token")
	}*/

	shop := router.Group("/shop") //endpoints for customers
	{
		shop.Use(middleware.UserIdentity(h.logger))

		categories := shop.Group("/categories")
		{
			categories.GET("/", h.ListCategories)

			products := categories.Group(":product_type")
			{
				products.GET("/", h.ListStandardizedProducts)
				products.GET("/:product_id", h.ShowProductInfo)
			}

		}

		adminPanel := shop.Group("/adminPanel") //endpoints for admins
		{
			adminPanel.Use(middleware.AdminOnly())

			categories := adminPanel.Group("/categories")
			{
				//categories.GET("/", h.ListCategoriesAdmin)
				//categories.POST("/", h.CreateCategory)
				//categories.PUT("/:id", h.UpdateCategory)
				//categories.DELETE("/:id", h.DeleteCategory)

				products := categories.Group(":category_name")
				{
					//products.GET("/", h.ListProductsAdmin)
					//products.GET("/:product_id", h.ShowProductInfoAdmin)
					products.POST("/", h.CreateProduct)
					//products.PATCH("/:product_id/updateGeneral", h.UpdateProductGeneral)
					//products.PUT("/:product_id/updateFull", h.UpdateProduct)
					products.DELETE("/:id", h.DeleteProduct)
				}
			}

			/*orders := shop.Group("/orders")
			{
				orders.GET("/")
				orders.GET("/:order_id")
				orders.POST("/")
				orders.PATCH("/:order_id")
				orders.DELETE("/:order_id")
			}*/

			/*users := router.Group("/users")
			{
				users.GET("/")
				users.GET("/:user_id")
				users.POST("/")
				users.PATCH("/")
				users.DELETE("/")
			}

			roles := router.Group("/roles")
			{
				roles.GET("/")
				roles.GET("/:role_id")
				roles.POST("/")
				roles.PATCH("/")
				roles.DELETE("/")
			}*/
		}
	}

	return router
}

func NewHandler(services *service.Service, logger *logging.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}
