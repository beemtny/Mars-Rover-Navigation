package models

import (
	"mars-rover-navigation/pkg"
)

type Obstacles struct {
	Obstacles []Position
}

func NewObstacles(obstacles []Position) Obstacles {
	return Obstacles{Obstacles: obstacles}
}

func (o Obstacles) IsCashObstacles(newPos Position) bool {
	for _, p := range o.Obstacles {
		if (newPos.X == p.X) && (newPos.Y == p.Y) {
			pkg.LogDebug("obstacle encountered!!")
			return true
		}
	}

	return false
}
