package main

// entry point for the server

import (
	"github.com/DylanCoon99/server-with-rate-limiter/src/server"
)



func main() {

	// create an instance of the server and start it here

	server.CreateServer("8080")
	server.Listen()

}