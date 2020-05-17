package main

import (
	"fmt"
	"os"
)

func getPort(defaultPort string) string {
	var port string = os.Getenv("PORT")

	// Set a default port if there is nothing in the environment
	if port == "" {
		port = defaultPort
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}