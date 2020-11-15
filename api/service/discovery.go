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
func GetLocation(distances ...float32) (x, y float32) {
	P1 := 0
	P2 := 1
	P3 := 2

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
			dist: distances[P1],
		},
		{
			name: "Skywalker",
			pLat: 100,
			pLon: -100,
			dist: distances[P2],
		},
		{
			name: "Sato",
			pLat: 500,
			pLon: 100,
			dist: distances[P3],
		},
	}

	xDist := math.Pow(float64(satPos[P1].pLon-satPos[P2].pLon), 2) + math.Pow(float64(satPos[P1].pLat-satPos[P2].pLat), 2)
	d := float32(math.Sqrt(xDist))

	ex := []float32{
		(satPos[P2].pLat - satPos[P1].pLat) / d,
		(satPos[P2].pLon - satPos[P1].pLon) / d,
	}

	p3p1 := []float32{
		satPos[P3].pLat - satPos[P1].pLat,
		satPos[P3].pLon - satPos[P1].pLon,
	}

	ival := ex[0]*p3p1[0] + ex[1] + p3p1[1]

	t1 := float64((satPos[P3].pLat - satPos[P1].pLat) - (ex[0] * ival))
	t2 := float64((satPos[P3].pLon - satPos[P1].pLon) - (ex[1] * ival))
	p3p1i := math.Pow(t1, 2) + math.Pow(t2, 2)

	ey := []float32{float32(t1 / math.Sqrt(p3p1i)), float32(t2 / math.Sqrt(p3p1i))}

	jval := (ey[0] * p3p1[0]) + (ey[1] * p3p1[1])

	xval := (math.Pow(float64(satPos[P1].dist), 2) - math.Pow(float64(satPos[P2].dist), 2)) + math.Pow(float64(d), 2)/float64((2*d))
	yval := ((math.Pow(float64(satPos[P1].dist), 2)-math.Pow(float64(satPos[P3].dist), 2))+math.Pow(float64(ival), 2)+math.Pow(float64(jval), 2))/
		float64((2*jval)) - (float64(ival/jval) * xval)

	posx := satPos[P1].pLat + (ex[0] * float32(xval)) + (ey[0] * float32(yval))
	posy := satPos[P1].pLon + (ex[1] * float32(xval)) + (ey[1] * float32(yval))

	return posx, posy
}
