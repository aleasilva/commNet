// Package service responsable to discovery the message position.
// Author : Alexandre.
package service

import "bytes"

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

//GetLocation =	Retorna as posicao do satelite, recebendo a posicao da mensagem.
func GetLocation(distances ...float32) (x, y float32) {

	satelitePosition := []struct {
		name string
		posX float32
		posY float32
	}{
		{
			name: "Kenobi",
			posX: -500,
			posY: -200,
		},
		{
			name: "Skywalker",
			posX: 100,
			posY: -100,
		},
		{
			name: "Sato",
			posX: 500,
			posY: 100,
		},
	}

	return satelitePosition[0].posX, satelitePosition[0].posX
}
