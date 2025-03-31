package pkg

import "fmt"

var debugMode = false

func LogDebug(msg string) {
	if !debugMode {
		return
	}

	fmt.Println(msg)
}
