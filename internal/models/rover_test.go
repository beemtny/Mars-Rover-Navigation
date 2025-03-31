package models

import (
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
