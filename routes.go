package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/vchicago/sso/controllers/v1"
)

func SetupRoutes(engine *gin.Engine) {
	engine.StaticFile("/", "./templates/docs.html")
	engine.StaticFile("/openapi.json", "./templates/openapi.json")

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PONG"})
	})

	v1Router := engine.Group("/v1")
	{
		v1Router.GET("/authorize", v1.GetAuthorize)
		v1Router.GET("/callback", v1.GetCallback)
		v1Router.GET("/certs", v1.GetCerts)
		v1Router.POST("/token", v1.PostToken)
	}
}
