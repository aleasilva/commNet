package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/russross/blackfriday"
)

type Satellites_stru struct {
	Satellites []struct {
		Name     string   `json:"name"`
		Distance float32  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellites"`
}

type DirectMessage struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

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

	router.POST("/topsecret", topSecretCall)
	router.POST("/topsecret_split/:satellite_name", topSecret_split)

	router.Run(":" + port)
}

func topSecretCall(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)

	var t Satellites_stru
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	log.Println(t.Satellites[1].Message[1])

	c.JSON(http.StatusOK, gin.H{"data": t.Satellites[0].Name})
}

func topSecret_split(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	satellite_name := strings.TrimPrefix(c.Request.URL.Path, "/topsecret_split/")

	fmt.Printf(satellite_name)

	var t DirectMessage
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	log.Println(t.Message[0])

	c.JSON(http.StatusOK, gin.H{"data": t.Distance})

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
	//Descobrir a maior matriz
	//Criar uma nova matriz para suportar as mensagens

	//Lendo os transmissores
	for i := 0; i < len(messages); i++ {
		//Lendo as mensagens
		for x := 0; x < len(messages[i]); x++ {
			buffer.WriteString(messages[i][x])
		}
	}

	return buffer.String()
}
