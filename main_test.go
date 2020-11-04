package main

import (
	"fmt"
	"testing"

	_ "github.com/heroku/x/hmetrics/onload"
)

func TestGetMessage(t *testing.T) {
	strMsg1 := []string{"este", "", "", "mensaje", ""}
	strMsg2 := []string{"", "es", "", "", "secreto"}
	strMsg3 := []string{"este", "", "un", "", ""}

	gotMsg := GetMessage(strMsg1, strMsg2, strMsg3)

	if gotMsg == "Ale" {
		t.Error("Ocorreu um erro!")
	}
	fmt.Println(gotMsg)

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

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMsg := GetMessage(tt.args.messages...); gotMsg != tt.wantMsg {
				t.Errorf("GetMessage() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
