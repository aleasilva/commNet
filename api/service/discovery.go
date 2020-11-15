// Package service responsable to discovery the message position.
// Author : Alexandre.
package service

import (
	"bytes"
	"math"
)

//GetMessage =	Read messages content
func GetMessage(messages ...[]string) (msg string) {
	var buffer bytes.Buffer
	var retMsg [5]string

	//Read the origin messages
	for indexMsg := 0; indexMsg < len(messages); indexMsg++ {

		for itemMsg := 0; itemMsg < len(messages[indexMsg]); itemMsg++ {
			if messages[indexMsg][itemMsg] != "" {
				retMsg[itemMsg] = messages[indexMsg][itemMsg]
			}

		}
	}

	//Prepare messagem to return
	for line := 0; line < len(retMsg); line++ {
		buffer.WriteString(retMsg[line] + " ")
	}

	return buffer.String()
}

//GetLocation =	Get the emissor position given the distance
//Reference = https://www.researchgate.net/post/Does-anyone-have-Trilateration-java-code
func GetLocation(distances ...float32) (x, y float32) {
	const Sindex1 = 0
	const Sindex2 = 1
	const Sindex3 = 2

	satPos := []struct {
		name string
		pLat float32
		pLon float32
		dist float32
	}{
		{
			name: "Kenobi",
			pLat: -500,
			pLon: -200,
			dist: distances[Sindex1],
		},
		{
			name: "Skywalker",
			pLat: 100,
			pLon: -100,
			dist: distances[Sindex2],
		},
		{
			name: "Sato",
			pLat: 500,
			pLon: 100,
			dist: distances[Sindex3],
		},
	}

	//Start the Trilateration Calc
	tmpDistAtoB := math.Pow(float64(satPos[Sindex1].pLon-satPos[Sindex2].pLon), 2) + math.Pow(float64(satPos[Sindex1].pLat-satPos[Sindex2].pLat), 2)
	distAtoB := float32(math.Sqrt(tmpDistAtoB))

	distAtoBCart := []float32{
		(satPos[Sindex2].pLat - satPos[Sindex1].pLat) / distAtoB,
		(satPos[Sindex2].pLon - satPos[Sindex1].pLon) / distAtoB,
	}

	distP3P1 := []float32{
		satPos[Sindex3].pLat - satPos[Sindex1].pLat,
		satPos[Sindex3].pLon - satPos[Sindex1].pLon,
	}

	ival := (distAtoBCart[0] * distP3P1[0]) + (distAtoBCart[1] * distP3P1[1])

	dist1 := float64((satPos[Sindex3].pLat - satPos[Sindex1].pLat) - (distAtoBCart[0] * ival))
	dist2 := float64((satPos[Sindex3].pLon - satPos[Sindex1].pLon) - (distAtoBCart[1] * ival))
	sumP3P1 := math.Pow(dist1, 2) + math.Pow(dist2, 2)

	distLong := []float32{float32(dist1 / math.Sqrt(sumP3P1)), float32(dist2 / math.Sqrt(sumP3P1))}

	sumDist := (distLong[0] * distP3P1[0]) + (distLong[1] * distP3P1[1])

	valX := (math.Pow(float64(satPos[Sindex1].dist), 2) - math.Pow(float64(satPos[Sindex2].dist), 2) +
		math.Pow(float64(distAtoB), 2)) / float64((2 * distAtoB))
	valY := ((math.Pow(float64(satPos[Sindex1].dist), 2)-math.Pow(float64(satPos[Sindex3].dist), 2))+
		math.Pow(float64(ival), 2)+math.Pow(float64(sumDist), 2))/float64((2*sumDist)) - (float64(ival/sumDist) * valX)

	posx := satPos[Sindex1].pLat + (distAtoBCart[0] * float32(valX)) + (distLong[0] * float32(valY))
	posy := satPos[Sindex1].pLon + (distAtoBCart[1] * float32(valX)) + (distLong[1] * float32(valY))

	return posx, posy
}
