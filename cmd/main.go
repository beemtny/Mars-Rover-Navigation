package main

import (
	"encoding/json"
	"fmt"
	"mars-rover-navigation/internal/app"
	"mars-rover-navigation/internal/models"
	"strings"
)

func main() {
	var (
		gridSize  int
		obstacles string
		commands  string
	)

	fmt.Print("Enter your grid_size (e.g., 5): ")
	_, err := fmt.Scan(&gridSize)
	if err != nil {
		fmt.Println("[Error] Error reading grid_size input:", err)
		return
	}

	fmt.Print("Enter your obstacles with no whitespace (e.g., [(1,2),(3,3)]): ")
	_, err = fmt.Scan(&obstacles)
	if err != nil {
		fmt.Println("[Error] Error reading obstacles input:", err)
		return
	}

	fmt.Print("Enter your commands (e.g., LMLMLMLMM): ")
	_, err = fmt.Scan(&commands)
	if err != nil {
		fmt.Println("[Error] Error reading obstacles input:", err)
		return
	}

	rover, status := app.NavigateRover(gridSize, parseObstacles(obstacles), strings.ToUpper(commands))
	resultByte, err := json.Marshal(models.Result{
		FinalPosition:  rover.CurrentPosition.ToFinalPosition(),
		FinalDirection: models.Directions[rover.DirectionsIndex],
		Status:         status,
	})
	if err != nil {
		fmt.Printf("[Error] Error marshalling result: %v", err)
		return
	}

	fmt.Println(string(resultByte))
}

func parseObstacles(obstacles string) []models.Position {
	var coords []models.Position

	obstacles = strings.Trim(obstacles, "[]")
	pairs := strings.Split(obstacles, "),(")

	for _, pair := range pairs {
		pair = strings.Trim(pair, "()")
		var x, y int
		_, err := fmt.Sscanf(pair, "%d,%d", &x, &y)
		if err == nil {
			coords = append(coords, models.Position{X: x, Y: y})
		}
	}

	return coords
}
