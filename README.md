# Mars Rover Navigation

## Project Description

This project implements a solution to navigate a rover on a grid-based planet, Mars. The rover is controlled by a series of commands that move it in different directions. The goal is to navigate the rover to a specific target position while avoiding obstacles and staying within the boundaries of the grid.

The solution implemented in this project is written in the Go programming language.The program takes input for the grid size, initial rover position, and direction, as well as a sequence of commands for the rover. It then processes these inputs and calculates the final rover position, direction, and status (whether it has reached the target position, out of bounderies or encountered an obstacle).

## Installation Instructions

1. Install Go (version 1.18 or later) from the official website: https://golang.org/dl/
2. Clone the project repository: `git clone https://github.com/beemtny/Mars-Rover-Navigation.git`
3. Navigate to the project directory: `cd Mars-Rover-Navigation`
4. Build the project: `go build -o mars_rover_navigation ./cmd/main.go`

## Usage Instructions

1. Run the program: `./mars_rover_navigation`
2. Enter the grid size (e.g., 5) followed by pressing Enter.
3. Enter the initial obstacles position (e.g., [(1,2),(3,3)]) followed by pressing Enter.
4. Enter the sequence of commands for the rover (e.g., LMLMLMLMM) followed by pressing Enter.
5. The program will output the final rover position, direction and status.

Example1:

input
```
grid_size = 5
obstacles = [(1,2),(3,3)]
commands = LMLMLMLMM
```

output (rover stops before moving out of bounds)
```
{
"final_position": [0,0],
"final_direction": "W",
"status": "Out of bounds"
}
```

Example2:

input
```
grid_size = 5
obstacles = [(1,2),(3,3)]
commands = LMLMLMLMMMM
```

output (rover stops before moving out of bounds)
```
{
"final_position": "[0,0]",
"final_direction": "W",
"status": "Out of bounds"
}
```

Example3:

input
```
grid_size = 5
obstacles = []
commands = MMMMMMMM
```

output (rover stops before moving out of bounds)
```
{
"final_position": "[0,4]",
"final_direction": "N",
"status": "Out of bounds"
}
```

Example4:

input
```
grid_size = 5
obstacles = []
commands = MRMR
```

output (success case)
```
{
"final_position": "[1,1]",
"final_direction": "S",
"status": "Success"
}
```

Example5:

input
```
grid_size = 5
obstacles = [(1,1)]
commands = MRMR
```

output (rover stops before cash obstacle)
```
{
"final_position": "[0,1]",
"final_direction": "E",
"status": "Obstacle encountered"
}
```


## Testing Instructions

1. Run the unit tests: `go test ./internal/... -v`
2. The test results will be displayed in the terminal.

## Additional Notes

1. The rover's movement commands (e.g., L, R, M) are case-sensitive. Ensure that the commands are entered in uppercase.
2. The input for the obstacles' positions should be provided without any whitespace. For example, instead of [(1, 2), (3, 3)], use [(1,2),(3,3)]. This ensures that the program can correctly parse and validate the obstacle positions.
3. If an obstacle is detected ahead, the rover must stop before moving and ignore the remaining commands.
4. If a rover gonna move out of boundaries, the rover must stop before moving and ignore the remaining commands.
