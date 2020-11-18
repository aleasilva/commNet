package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/api/service"
)

//DirectMessage = Save messages
type DirectMessage struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
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
func TopSecretSplit(context *gin.Context) {

	decoder := json.NewDecoder(context.Request.Body)
	satelliteName := strings.TrimPrefix(context.Request.URL.Path, "/topsecret_split/")

	var directMsg DirectMessage
	err := decoder.Decode(&directMsg)

	if err != nil {
		panic(err)
	}

	msgReturn := service.GetMessage(directMsg.Message)

	//Recover the position of the sattlelite
	//locX, locY := service.GetLocation(directMsg.Distance, directMsg.Distance, directMsg.Distance)

	//Prepare response
	position := position{
		X: 0,
		Y: 0,
	}
	position.X, position.Y = getSattelitePositionFromName(satelliteName)

	if position.X == 0 && position.Y == 0 {
		context.JSON(http.StatusBadRequest, "")
	}

	response := topSecretResponse{
		Message:  msgReturn,
		Position: &position,
	}

	context.JSON(http.StatusOK, response)
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

	response := topSecretResponse{
		Message:  msgReturn,
		Position: &position,
	}

	context.JSON(http.StatusOK, response)

}

//getDistanceInOrder Recover the distance from each sattelite based on name
func getDistanceInOrder(message MessageProtocolStru) (dist01, dist02, dist03 float32) {

	for satIndex := 0; satIndex < len(message.Satellites); satIndex++ {
		satName := strings.ToUpper(message.Satellites[satIndex].Name)
		satDist := message.Satellites[satIndex].Distance

		if satName == "KENOBI" {
			dist01 = satDist
		} else if satName == "SKYWALKER" {
			dist02 = satDist
		} else if satName == "SATO" {
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

//getSattelitePositionFromName = Retorna a posicao de um satellite baseado no nome
func getSattelitePositionFromName(name string) (x, y float32) {

	satPos := []struct {
		name string
		pLat float32
		pLon float32
		dist float32
	}{
		{
			name: "KENOBI",
			pLat: -500,
			pLon: -200,
		},
		{
			name: "SKYWALKER",
			pLat: 100,
			pLon: -100,
		},
		{
			name: "SATO",
			pLat: 500,
			pLon: 100,
		},
	}

	x = 0
	y = 0

	for _, v := range satPos {
		if v.name == strings.ToUpper(name) {
			x = v.pLat
			y = v.pLon
		}
	}

	return x, y

}
