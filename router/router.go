package router

import (
	"github.com/gin-gonic/gin"
	"health/api"
	"health/middlewares"
)

func InitRouter(r *gin.Engine) {
	r.Use(middlewares.Cors())
	r.GET("/version", api.Version)

	user := r.Group("/user")
	{
		user.POST("/login", api.Login)
		user.POST("/register", api.Register)
	}

	record := r.Group("/record").Use(middlewares.Auth())
	{
		record.POST("/create", api.Create)
		record.POST("/list", api.List)
		record.POST("/del", api.Del)
	}
}
