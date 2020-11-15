// Copyright inSSoft Corp.
// All Rights Reserved
//
// Program responsable for process the requests from
// our bases around the space.
// Author : Alexandre.

package main

import (
	"log"
	"os"

	"github.com/heroku/go-getting-started/api/controller"
	_ "github.com/heroku/x/hmetrics/onload"
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

	router := controller.SetupServer()

	router.Run(":" + port)
}
