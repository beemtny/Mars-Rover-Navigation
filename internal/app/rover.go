package app

import (
	"fmt"
	"mars-rover-navigation/internal/models"
	"mars-rover-navigation/pkg"
)

func NavigateRover(gridSize int, obs []models.Position, commands string) (models.Rover, models.Status) {
	var (
		rover     = models.NewRover()
		grid      = models.NewGrid(gridSize)
		obstacles = models.NewObstacles(obs)
	)

	for _, c := range commands {
		cmd := string(c)
		pkg.LogDebug(fmt.Sprintf("Commnad: %s ,Current direction: %s, current position: %v\n", cmd, models.Directions[rover.DirectionsIndex], rover.CurrentPosition))

		switch cmd {
		case "L", "R":
			rover.DirectionsIndex = rover.Rotate(cmd)
		case "M":
			newPos := rover.MoveForward()

			// chank boundaries
			if isOOB := grid.IsOutOfBounderies(newPos); isOOB {
				return rover, models.OOB
			}

			// check obstacles
			if isCash := obstacles.IsCashObstacles(rover.CurrentPosition); isCash {
				return rover, models.ObstacleEncountered
			}

			rover.CurrentPosition = newPos
		default:
			pkg.LogDebug(fmt.Sprintf("invalid command: %s", cmd))
		}

		pkg.LogDebug(fmt.Sprintf("New direction: %s, current position: %v", models.Directions[rover.DirectionsIndex], rover.CurrentPosition))
		pkg.LogDebug("--------------------------------")
	}

	return rover, models.Success
}
