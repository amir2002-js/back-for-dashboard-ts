package routers

import (
	"github.com/gin-gonic/gin"
	"paysee2/handlers"
	"paysee2/middlewares"
)

func Router(r *gin.Engine, handlers *handlers.Handlers) {
	api := r.Group("/api")
	{
		api.POST("/refreshToken", handlers.AccessTokenHandler)
		register := api.Group("/register")
		{
			register.POST("/", handlers.UserHandler.RegisterHandler)
		}

		login := api.Group("/login")
		{
			login.POST("/", handlers.UserHandler.LoginHandler)
		}

		admin := api.Group("/admin")
		admin.Use(middlewares.CheckWho(handlers.DB))
		{
			admin.GET("/users", handlers.UserHandler.GetAllUsersHandler)
		}

		simpleUser := api.Group("/simpleUser")
		simpleUser.Use(middlewares.CheckWho(handlers.DB))
		{
			customers := simpleUser.Group("/customers")
			{
				customers.GET("/:customerType", handlers.CustomerHandler.GetCustomerByTypeHandler)
				customers.GET("/:id", handlers.CustomerHandler.GetCustomerByIdHandler)
				customers.POST("/", handlers.CustomerHandler.CreateCustomerHandler)
			}
		}
	}
}
