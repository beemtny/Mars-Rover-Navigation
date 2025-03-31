package models

type Status string

const (
	Success             Status = "Success"
	ObstacleEncountered Status = "Obstacle encountered"
	OOB                 Status = "Out of bounds"
)

type Result struct {
	FinalPosition  string `json:"final_position"`
	FinalDirection string `json:"final_direction"`
	Status         Status `json:"status"`
}
