// Copyright inSSoft Corp.
// All Rights Reserved
//
// Program responsable for process the requests from
// our bases around the space.
// Author : Alexandre.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/api/controller"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set")
		port = "5000"
		//log.Fatal("$PORT must be set")
	}

	tStr := os.Getenv("REPEAT")
	if tStr == "" {
		tStr = "2"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/mark", func(c *gin.Context) {
		c.String(http.StatusOK, string(blackfriday.Run([]byte("**hi!**"))))
	})

	//Router for the service requested.
	router.POST("/topsecret", controller.TopSecretCall)
	router.POST("/topsecret_split/:satellite_name", controller.TopSecretSplit)

	router.Run(":" + port)
}
