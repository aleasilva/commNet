// Copyright inSSoft Corp.
// All Rights Reserved
//
// Classe responsavel por receber as requisições para realizar o processamento.
// Author : Alexandre.

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

//SatellitesStru = Estrutura para realizar o parser do json recebido para processamento.
type SatellitesStru struct {
	Satellites []struct {
		Name     string   `json:"name"`
		Distance float32  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellites"`
}

//DirectMessage = Estrutura para armazenamento das mensagens
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
	router.POST("/topsecret_split/:satellite_name", topSecretSplit)

	router.Run(":" + port)
}

func topSecretCall(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)

	var t SatellitesStru
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	log.Println(t.Satellites[1].Message[1])

	c.JSON(http.StatusOK, gin.H{"data": t.Satellites[0].Name})
}

//topSecretSplit = Faz a separação da mensagem recebida
func topSecretSplit(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	satelliteName := strings.TrimPrefix(c.Request.URL.Path, "/topsecret_split/")

	fmt.Printf(satelliteName)

	var t DirectMessage
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	log.Println(t.Message[0])

	c.JSON(http.StatusOK, gin.H{"data": t.Distance})

}

//GetLocation =	Retorna as posicao do satelite, recebendo a posicao da mensagem
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

//GetMessage =	Lendo o conteudo da mensagem
func GetMessage(messages ...[]string) (msg string) {
	var buffer bytes.Buffer
	var retMsg [5]string

	//Descobrir a maior matriz
	//Criar uma nova matriz para suportar as mensagens
	retMsg[0] = "Alexandre"

	//Lendo os transmissores
	for i := 0; i < len(messages); i++ {

		//Lendo as mensagens
		for x := 0; x < len(messages[i]); x++ {
			if messages[i][x] != "" {
				retMsg[x] = messages[i][x]
			}

		}
	}

	//Montado a mensagem para retorno
	for line := 0; line < len(retMsg); line++ {
		buffer.WriteString(retMsg[line] + " ")
	}

	return buffer.String()
}
