package models

import (
	"reflect"
	"testing"
)

func TestRover_Rotate(t *testing.T) {
	type rover struct {
		DirectionsIndex int
	}
	type args struct {
		direction string
	}
	tests := []struct {
		name                  string
		fields                rover
		args                  args
		wantNewDirectionIndex int
	}{
		{
			name: "start direction: N rotate direction: Left",
			fields: rover{
				DirectionsIndex: 0,
			},
			args: args{
				direction: "L",
			},
			wantNewDirectionIndex: 3,
		},
		{
			name: "start direction: N rotate direction: Right",
			fields: rover{
				DirectionsIndex: 0,
			},
			args: args{
				direction: "R",
			},
			wantNewDirectionIndex: 1,
		},
		{
			name: "start direction: E rotate direction: Left",
			fields: rover{
				DirectionsIndex: 1,
			},
			args: args{
				direction: "L",
			},
			wantNewDirectionIndex: 0,
		},
		{
			name: "start direction: E rotate direction: Right",
			fields: rover{
				DirectionsIndex: 1,
			},
			args: args{
				direction: "R",
			},
			wantNewDirectionIndex: 2,
		},
		{
			name: "start direction: S rotate direction: Left",
			fields: rover{
				DirectionsIndex: 2,
			},
			args: args{
				direction: "L",
			},
			wantNewDirectionIndex: 1,
		},
		{
			name: "start direction: S rotate direction: Right",
			fields: rover{
				DirectionsIndex: 2,
			},
			args: args{
				direction: "R",
			},
			wantNewDirectionIndex: 3,
		},
		{
			name: "start direction: W rotate direction: Left",
			fields: rover{
				DirectionsIndex: 3,
			},
			args: args{
				direction: "L",
			},
			wantNewDirectionIndex: 2,
		},
		{
			name: "start direction: W rotate direction: Right",
			fields: rover{
				DirectionsIndex: 3,
			},
			args: args{
				direction: "R",
			},
			wantNewDirectionIndex: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rover{
				DirectionsIndex: tt.fields.DirectionsIndex,
			}
			if gotNewDirectionIndex := r.Rotate(tt.args.direction); gotNewDirectionIndex != tt.wantNewDirectionIndex {
				t.Errorf("Rover.Rotate() = %v, want %v", gotNewDirectionIndex, tt.wantNewDirectionIndex)
			}
		})
	}
}

func TestRover_MoveForward(t *testing.T) {
	type rover struct {
		CurrentPosition Position
		DirectionsIndex int
	}
	tests := []struct {
		name       string
		fields     rover
		wantNewPos Position
	}{
		{
			name: "move forward direction North",
			fields: rover{
				CurrentPosition: Position{X: 2, Y: 2},
				DirectionsIndex: 0,
			},
			wantNewPos: Position{X: 2, Y: 3},
		},
		{
			name: "move forward direction East",
			fields: rover{
				CurrentPosition: Position{X: 2, Y: 2},
				DirectionsIndex: 1,
			},
			wantNewPos: Position{X: 3, Y: 2},
		},
		{
			name: "move forward direction South",
			fields: rover{
				CurrentPosition: Position{X: 2, Y: 2},
				DirectionsIndex: 2,
			},
			wantNewPos: Position{X: 2, Y: 1},
		},
		{
			name: "move forward direction West",
			fields: rover{
				CurrentPosition: Position{X: 2, Y: 2},
				DirectionsIndex: 3,
			},
			wantNewPos: Position{X: 1, Y: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rover{
				CurrentPosition: tt.fields.CurrentPosition,
				DirectionsIndex: tt.fields.DirectionsIndex,
			}
			if gotNewPos := r.MoveForward(); !reflect.DeepEqual(gotNewPos, tt.wantNewPos) {
				t.Errorf("Rover.MoveForward() = %v, want %v", gotNewPos, tt.wantNewPos)
			}
		})
	}
}
