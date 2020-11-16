package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/api/service"
)

//DirectMessage = Save messages
type DirectMessage struct {
	Distance float32  `json:"distance"`
	Messages []string `json:"message"`
}

//MessageProtocolStru = Used to parse json request with the message
// to be verified
type MessageProtocolStru struct {
	Satellites []struct {
		Name     string   `json:"name"`
		Distance float32  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellites"`
}

type topSecretResponse struct {
	Message  string
	Position *position
}

type position struct {
	X float32
	Y float32
}

//TopSecretSplit = Split the received messages
func TopSecretSplit(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)
	satelliteName := strings.TrimPrefix(c.Request.URL.Path, "/topsecret_split/")
	fmt.Printf(satelliteName)

	var t DirectMessage
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	log.Println(t.Messages[0])

	c.JSON(http.StatusOK, gin.H{"dataNew": t.Distance})
}

//TopSecretCall = Process secrets Call from ours ships
func TopSecretCall(context *gin.Context) {
	decoder := json.NewDecoder(context.Request.Body)

	var message MessageProtocolStru
	decoderResult := decoder.Decode(&message)

	if decoderResult != nil {
		panic(decoderResult)
	}

	//Recover the message
	msgReturn := service.GetMessage(message.Satellites[0].Message,
		message.Satellites[1].Message,
		message.Satellites[2].Message)

	//Recover the orgin of the message
	locX, locY := service.GetLocation(getDistanceInOrder(message))

	//Prepare response
	position := position{
		X: locX,
		Y: locY,
	}

	reponse := topSecretResponse{
		Message:  msgReturn,
		Position: &position,
	}

	context.JSON(http.StatusOK, reponse)

}

//getDistanceInOrder Recover the distance from each sattelite based on name
func getDistanceInOrder(message MessageProtocolStru) (dist01, dist02, dist03 float32) {

	for satIndex := 0; satIndex < len(message.Satellites); satIndex++ {
		satName := message.Satellites[satIndex].Name
		satDist := message.Satellites[satIndex].Distance

		if satName == "kenobi" {
			dist01 = satDist
		} else if satName == "skywalker" {
			dist02 = satDist
		} else if satName == "sato" {
			dist03 = satDist
		}
	}
	return dist01, dist02, dist03
}

//PingEndpoint = Response I am alive, for the client.
func PingEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
