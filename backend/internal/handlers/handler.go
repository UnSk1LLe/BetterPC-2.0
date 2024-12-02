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
	"github.com/patrickmn/go-cache"
)

type Handler struct {
	services *service.Service
	logger   *logging.Logger
	cfg      *configs.Config
	cache    *cache.Cache
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(gin.Logger())

	sessions.Init(router)

	router.SetFuncMap(templateFunctions.TmplFuncs)
	html.LoadTemplates(router)
	static.LoadStatic(router)

	middleware := middlewares.NewMiddleware(h.services, h.logger, h.cfg, h.cache)

	api := router.Group("/api")

	apiV1 := api.Group("/v1")

	auth := apiV1.Group("/auth")
	{
		auth.Use(middleware.RateLimitFromClient(middlewares.DefaultRateLimit, middlewares.DefaultRateInterval))

		//auth.GET("/register", h.RegisterForm)
		auth.POST("/register", h.Register)
		//auth.GET("/login", h.LoginForm)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
		auth.POST("/logout", h.Logout)
	}

	verification := apiV1.Group("/verification")
	{
		verification.Use(middleware.RateLimitFromClient(middlewares.DefaultRateLimit, middlewares.DefaultRateInterval))

		verification.GET("/", h.SendVerificationLink)
		verification.POST("/:token", h.VerifyUser)
	}

	passwordRecovery := apiV1.Group("/password_recovery")
	{
		verification.Use(middleware.RateLimitFromClient(middlewares.DefaultRateLimit, middlewares.DefaultRateInterval))

		passwordRecovery.GET("/", h.SendRecoveryLink)
		passwordRecovery.POST("/:token", h.RecoverPassword)
	}

	shop := apiV1.Group("/shop") //endpoints for customers
	{
		shop.Use(middleware.UserIdentity())

		categories := shop.Group("/categories")
		{
			categories.GET("/", h.ListCategories)

			products := categories.Group(":product_type")
			{
				products.GET("/", h.ListStandardizedProducts)
				products.GET("/:id", h.ShowProductInfo)
			}
		}

		orders := shop.Group("/orders")
		{
			orders.Use(middleware.IsAuthorized())

			orders.GET("/", h.ListUserOrders)
			orders.GET("/:id", h.GetUserOrder)
			orders.POST("/", h.CreateOrderWithItemHeaders)
			orders.POST("/:id/payment", h.ProcessOrderPayment)
			//orders.PATCH("/:id/update", h.UpdateUserOrder)
			orders.PATCH("/:id/cancel", h.CancelUserOrder)
		}

		userInfo := shop.Group("/user_info")
		{
			userInfo.Use(middleware.IsAuthorized())

			userInfo.GET("/", h.GetUserInfo)
			userInfo.PATCH("/", h.UpdateUserInfo)
			userInfo.POST("/verification", h.SendNewVerificationLink)
			userInfo.POST("/image", h.UploadUserImage)

			paymentDetails := userInfo.Group("/payment_methods")
			{
				paymentDetails.GET("/", h.GetUserPaymentMethods)
				paymentDetails.POST("/:id", h.AttachNewPaymentMethod)
				paymentDetails.DELETE("/:id", h.RemovePaymentMethod)
			}
		}
	}

	adminPanel := apiV1.Group("/admin_panel") //endpoints for admins
	{
		adminPanel.Use(middleware.UserIdentity(), middleware.StaffOnly())

		categories := adminPanel.Group("/categories")
		{
			//categories.GET("/", h.ListCategoriesAdmin)
			//categories.POST("/", h.CreateCategory)
			//categories.PUT("/:id", h.UpdateCategory)
			//categories.DELETE("/:id", h.DeleteCategory)

			products := categories.Group(":product_type")
			{
				products.GET("/", h.ListStandardizedProducts)
				products.GET("/:id", h.ShowProductInfo)
				products.POST("/", h.CreateProduct)
				products.PATCH("/:id/update_general", h.UpdateProductGeneral)
				products.PATCH("/:id/update_details", h.UpdateProduct)
				products.DELETE("/:id", middleware.AdminOnly(), h.DeleteProduct)
			}
		}

		orders := adminPanel.Group("/orders")
		{
			orders.GET("/", h.ListOrders)
			orders.GET("/:order_id", h.OrderDetails)
			orders.POST("/", h.CreateOrderWithItemHeaders)
			orders.PATCH("/:id/cancel", h.CancelOrder)
			//orders.PATCH("/:order_id", h.UpdateOrder)
			orders.DELETE("/:id", middleware.AdminOnly(), h.DeleteOrder)
		}

		users := adminPanel.Group("/users")
		{
			users.Use(middleware.AdminOnly())

			users.GET("/", h.GetUserList)
			users.GET("/:id", h.GetUser)
			users.POST("/", h.CreateUser)
			users.PATCH("/", h.UpdateUser)
			users.DELETE("/", h.DeleteUser)
		}
	}

	return router
}

func NewHandler(services *service.Service, logger *logging.Logger, cfg *configs.Config, lCache *cache.Cache) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
		cfg:      cfg,
		cache:    lCache,
	}
}
