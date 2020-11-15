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
	//Recover port from OS.
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
		log.Println("$PORT must be set, on os system, check heroku to see how.")
	}

	router := controller.SetupServer()

	router.Run(":" + port)
}
