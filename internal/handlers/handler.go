package handlers

import (
	"BetterPC_2.0/internal/service"
	"BetterPC_2.0/pkg/html"
	"BetterPC_2.0/pkg/logging"
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

	router.SetFuncMap(templateFunctions.TmplFuncs)
	html.LoadTemplates(router)
	static.LoadStatic(router)

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

		categories := shop.Group("/categories", h.UserIdentity)
		{
			categories.GET("/", h.ListCategories)

			products := categories.Group(":product_type/products")
			{
				products.GET("/", h.ListStandardizedProducts)
				products.GET("/:product_id", h.ShowProductInfo)
			}

		}

		adminPanel := shop.Group("/adminPanel") //endpoints for admins
		{
			categories := adminPanel.Group("/categories")
			{
				//categories.GET("/", h.ListCategoriesAdmin)
				//categories.POST("/", h.CreateCategory)
				//categories.PUT("/:id", h.UpdateCategory)
				//categories.DELETE("/:id", h.DeleteCategory)

				products := categories.Group(":category_name/products")
				{
					//products.GET("/:product_id", h.ShowProductInfoAdmin)
					products.POST("/", h.CreateProduct)
					//products.PUT("/:product_id/updateGeneral", h.UpdateProductGeneral)
					//products.PUT("/:product_id/updateFull", h.UpdateProduct)
					products.DELETE("/:id", h.DeleteProduct)
				}
			}

			/*orders := shop.Group("/orders")
			{
				orders.GET("/")
				orders.GET("/:order_id")
				orders.POST("/")
				orders.PUT("/:order_id")
				orders.DELETE("/:order_id")
			}*/

			/*users := router.Group("/users")
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
