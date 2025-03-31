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

func (o Obstacles) IsCashObstacles(currPos Position) bool {
	for _, p := range o.Obstacles {
		if (currPos.X == p.X) && (currPos.Y == p.Y) {
			pkg.LogDebug("obstacle encountered!!")
			return true
		}
	}

	return false
}
