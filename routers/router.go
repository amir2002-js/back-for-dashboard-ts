package routers

import (
	"github.com/gin-gonic/gin"
	"paysee2/handlers/userHandlers"
)

func Router(r *gin.Engine, userHandlers *userHandlers.UserHandlers) {
	api := r.Group("/api")
	{
		register := api.Group("/register")
		{
			register.POST("/", userHandlers.CreateUser)
			register.POST("/verify", nil)
		}

		login := api.Group("/login")
		{
			login.POST("/", nil)
			login.POST("/verify", nil)
		}

		kyc := api.Group("/kyc")
		{
			kyc.POST("/", nil)
			kyc.POST("/status", nil)
		}

		api.POST("/reset-password", nil)
	}
}
