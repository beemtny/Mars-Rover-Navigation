package models

import (
	"fmt"
	"mars-rover-navigation/pkg"
)

type Rover struct {
	CurrentPosition Position
	DirectionsIndex int
}

func NewRover() Rover {
	return Rover{
		CurrentPosition: Position{X: 0, Y: 0},
		DirectionsIndex: 0, // North
	}
}

type Position struct {
	X int
	Y int
}

func (p Position) ToFinalPosition() string {
	return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

var Directions = []string{"N", "E", "S", "W"}
var Moves = map[string]Position{
	"N": {0, 1},
	"E": {1, 0},
	"S": {0, -1},
	"W": {-1, 0},
}

func (r Rover) Rotate(direction string) (newDirectionIndex int) {
	switch direction {
	case "L":
		pkg.LogDebug("turn left")
		newDirectionIndex = (r.DirectionsIndex + 3) % 4
	case "R":
		pkg.LogDebug("turn right")
		newDirectionIndex = (r.DirectionsIndex + 1) % 4
	default:
		pkg.LogDebug("invalid direction")
	}

	return newDirectionIndex
}

func (r Rover) MoveForward() (newPos Position) {
	pkg.LogDebug("move forward")
	direction := Directions[r.DirectionsIndex]
	newX := r.CurrentPosition.X + Moves[direction].X
	newY := r.CurrentPosition.Y + Moves[direction].Y

	return Position{X: newX, Y: newY}
}
