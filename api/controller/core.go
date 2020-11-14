package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//DirectMessage = Estrutura para armazenamento das mensagens
type DirectMessage struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

//SatellitesStru = Estrutura para realizar o parser do json recebido para processamento.
type SatellitesStru struct {
	Satellites []struct {
		Name     string   `json:"name"`
		Distance float32  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellites"`
}

//TopSecretSplit = Faz a separação da mensagem recebida
func TopSecretSplit(c *gin.Context) {

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

//TopSecretCall = Process secrets Call from ours ships
func TopSecretCall(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)

	var t SatellitesStru
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	log.Println(t.Satellites[1].Message[1])

	c.JSON(http.StatusOK, gin.H{"dataNew": t.Satellites[0].Name})
}
