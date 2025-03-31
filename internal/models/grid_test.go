package models

import "testing"

func TestGrid_IsOutOfBounderies(t *testing.T) {
	type grid struct {
		Size int
	}
	type args struct {
		newPos Position
	}
	tests := []struct {
		name   string
		fields grid
		args   args
		want   bool
	}{
		{
			name:   "X out of bounderies minus side",
			fields: grid{Size: 5},
			args:   args{newPos: Position{X: -1, Y: 0}},
			want:   true,
		},
		{
			name:   "X out of bounderies over grid size",
			fields: grid{Size: 5},
			args:   args{newPos: Position{X: 5, Y: 0}},
			want:   true,
		},
		{
			name:   "Y out of bounderies minus side",
			fields: grid{Size: 5},
			args:   args{newPos: Position{X: 4, Y: -1}},
			want:   true,
		},
		{
			name:   "Y out of bounderies over grid size",
			fields: grid{Size: 5},
			args:   args{newPos: Position{X: 2, Y: 6}},
			want:   true,
		},
		{
			name:   "position in bounderies",
			fields: grid{Size: 5},
			args:   args{newPos: Position{X: 2, Y: 2}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Grid{
				Size: tt.fields.Size,
			}
			if got := g.IsOutOfBounderies(tt.args.newPos); got != tt.want {
				t.Errorf("Grid.IsOutOfBounderies() = %v, want %v", got, tt.want)
			}
		})
	}
}
