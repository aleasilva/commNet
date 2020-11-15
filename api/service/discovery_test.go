// Copyright inSSoft Corp.
// All Rights Reserved
//
// Service responsable to discovery the message position.
// Author : Alexandre.
package service

import "testing"

func TestGetLocation(t *testing.T) {
	type args struct {
		distances []float32
	}
	tests := []struct {
		name  string
		args  args
		wantX float32
		wantY float32
	}{
		{
			name:  "Test Get Location",
			args:  args{distances: []float32{100, 115.5, 142.7}},
			wantX: 2158.0815,
			wantY: -2558.356,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := GetLocation(tt.args.distances...)
			if gotX != tt.wantX {
				t.Errorf("GetLocation() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("GetLocation() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestGetMessage(t *testing.T) {
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

func TestGetMessageOld(t *testing.T) {
	strMsg1 := []string{"este", "", "", "mensaje", ""}
	strMsg2 := []string{"", "es", "", "", "secreto"}
	strMsg3 := []string{"este", "", "un", "", ""}

	gotMsg := GetMessage(strMsg1, strMsg2, strMsg3)

	if gotMsg != "este es un mensaje secreto " {
		t.Error("Ocorreu um erro no método GetMessage...")
	}

}
