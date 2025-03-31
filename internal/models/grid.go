package models

import "mars-rover-navigation/pkg"

type Grid struct {
	Size int
}

func NewGrid(size int) Grid {
	return Grid{Size: size}
}

func (g Grid) IsOutOfBounderies(newPos Position) bool {
	if newPos.X < 0 || newPos.X >= g.Size || newPos.Y < 0 || newPos.Y >= g.Size {
		pkg.LogDebug("cannot move forward, cause out of bounderies position")
		return true
	}

	return false
}
