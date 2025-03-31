package app

import (
	"mars-rover-navigation/internal/models"
	"reflect"
	"testing"
)

func TestNavigateRover(t *testing.T) {
	type args struct {
		gridSize  int
		obstacles []models.Position
		commands  string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Rover
		want1   models.Status
		wantErr bool
	}{
		{
			name: "Normal movements without obstacles",
			args: args{
				gridSize:  5,
				obstacles: []models.Position{},
				commands:  "MMMM",
			},
			want: models.Rover{
				CurrentPosition: models.Position{X: 0, Y: 4},
				DirectionsIndex: 0,
			},
			want1:   models.Success,
			wantErr: false,
		},
		{
			name: "Collisions with obstacles in different directions",
			args: args{
				gridSize:  5,
				obstacles: []models.Position{{X: 1, Y: 1}},
				commands:  "MRM",
			},
			want: models.Rover{
				CurrentPosition: models.Position{X: 0, Y: 1},
				DirectionsIndex: 1,
			},
			want1:   models.ObstacleEncountered,
			wantErr: false,
		},
		{
			name: "Attempts to move out of grid boundaries",
			args: args{
				gridSize:  2,
				obstacles: []models.Position{},
				commands:  "MMMM",
			},
			want: models.Rover{
				CurrentPosition: models.Position{X: 0, Y: 1},
				DirectionsIndex: 0,
			},
			want1:   models.OOB,
			wantErr: false,
		},
		{
			name: "Invalid commands",
			args: args{
				gridSize:  2,
				obstacles: []models.Position{},
				commands:  "MDDD",
			},
			want:    models.Rover{},
			want1:   "",
			wantErr: true,
		},
		{
			name: "Minimal grid sizes",
			args: args{
				gridSize:  1,
				obstacles: []models.Position{},
				commands:  "M",
			},
			want: models.Rover{
				CurrentPosition: models.Position{X: 0, Y: 0},
				DirectionsIndex: 0,
			},
			want1:   models.OOB,
			wantErr: false,
		},
		{
			name: "Large grid sizes",
			args: args{
				gridSize:  10000,
				obstacles: []models.Position{},
				commands:  "MMMMMMMMMMMMMMRMMRMLMMLMRM",
			},
			want: models.Rover{
				CurrentPosition: models.Position{X: 5, Y: 14},
				DirectionsIndex: 1,
			},
			want1:   models.Success,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := NavigateRover(tt.args.gridSize, tt.args.obstacles, tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("NavigateRover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NavigateRover() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NavigateRover() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
