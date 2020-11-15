package controller

import (
	"github.com/gin-gonic/gin"
)

//SetupServer = The engine with all endpoints is now extracted from the main function
func SetupServer() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())

	//router := gin.Default()

	router.GET("/ping", PingEndpoint)

	router.POST("/topsecret", TopSecretCall)

	router.POST("/topsecret_split/:satellite_name", TopSecretSplit)

	return router
}
