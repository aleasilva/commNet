package main

import (
	"testing"

	_ "github.com/heroku/x/hmetrics/onload"
)

func TestGetMessage(t *testing.T) {
	strMsg1 := []string{"este", "", "", "mensaje", ""}
	strMsg2 := []string{"", "es", "", "", "secreto"}
	strMsg3 := []string{"este", "", "un", "", ""}

	gotMsg := GetMessage(strMsg1, strMsg2, strMsg3)

	if gotMsg != "este es un mensaje secreto " {
		t.Error("Ocorreu um erro no método GetMessage...")
	}

}

func TestGetMessage2(t *testing.T) {
	type args struct {
		messages [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantMsg string
	}{
		{
			name: "Teste 1",
			args: args{messages: [][]string{
				{"este", "", "", "mensaje", ""},
				{"", "es", "", "", "secreto"},
				{"este", "", "un", "", ""},
			}},
			wantMsg: "este es un mensaje secreto ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMsg := GetMessage(tt.args.messages...); gotMsg != tt.wantMsg {
				t.Errorf("GetMessage() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
