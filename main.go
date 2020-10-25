package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tStr := os.Getenv("REPEAT")
	repeat, err := strconv.Atoi(tStr)
	if err != nil {
		log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
		repeat = 5
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

	router.GET("/repeat", repeatHandler(repeat))
	router.POST("/upload", repeatHandler(repeat))
	router.POST("/topsecret", tratarTreta())

	router.Run(":" + port)
}

func tratarTreta() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buffer bytes.Buffer

		buffer.WriteString("Tratar Treta")

		c.String(http.StatusOK, buffer.String())
	}

}

func repeatHandler(r int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buffer bytes.Buffer
		for i := 0; i < r; i++ {
			buffer.WriteString("Hello from Go!\n")
		}
		c.String(http.StatusOK, buffer.String())
	}

}

/*
	Retorna as posicao do satelite, recebendo a posicao da mensagem
*/
func GetLocation(distances ...float32) (x, y float32) {

	if distances[0] > -499 && distances[1] < 201 {
		return -500, -200
	} else if distances[0] > -99 && distances[1] < 101 {
		return -100, 100
	} else if distances[0] > 99 && distances[1] < 499 {
		return 100, 500
	} else {
		return 0, 0
	}

}

/*
	Lendo o conteudo da mensagem
*/
func GetMessage(messages ...[]string) (msg string) {
	var buffer bytes.Buffer

	//Lendo os transmissores
	for i := 0; i < len(messages); i++ {
		//Lendo as mensagens
		for x := 0; x < len(messages[i]); x++ {
			buffer.WriteString(messages[x][i])
		}
	}

	return buffer.String()
}
