package main

import (
	"fmt"
	"os"
	"strconv"
	"webserver/tutorial/webserver"
)

func main() {
	serverHost := os.Getenv("SERVER_HOST")
	if serverHost == "" {
		fmt.Println("No SERVER_HOST environment variable set")
		return
	}

	serverPortStr := os.Getenv("SERVER_PORT")
	if serverPortStr == "" {
		fmt.Println("No SERVER_PORT environment variable set")
		return
	}

	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		fmt.Printf("Failed to convert SERVER_PORT environment variable to integer: %v\n", err)
		return
	}

	fmt.Printf("Starting server listening on %s:%d ...\n", serverHost, serverPort)
	if err := webserver.Listen(serverHost, serverPort); err != nil {
		fmt.Printf("Error while running server: %v\n", err)
		return
	}
}
