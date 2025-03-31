package models

import "testing"

func TestObstacles_IsCashObstacles(t *testing.T) {
	type fields struct {
		Obstacles []Position
	}
	type args struct {
		currPos Position
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "found obstacle",
			fields: fields{
				Obstacles: []Position{{X: 2, Y: 2}, {X: 1, Y: 1}},
			},
			args: args{
				currPos: Position{X: 1, Y: 1},
			},
			want: true,
		},
		{
			name: "not found obstacle",
			fields: fields{
				Obstacles: []Position{{X: 2, Y: 2}, {X: 3, Y: 3}},
			},
			args: args{
				currPos: Position{X: 1, Y: 1},
			},
			want: false,
		},
		{
			name: "empty obstacle",
			fields: fields{
				Obstacles: []Position{},
			},
			args: args{
				currPos: Position{X: 1, Y: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Obstacles{
				Obstacles: tt.fields.Obstacles,
			}
			if got := o.IsCashObstacles(tt.args.currPos); got != tt.want {
				t.Errorf("Obstacles.IsCashObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
