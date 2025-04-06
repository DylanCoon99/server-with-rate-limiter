package main

// entry point for the client

import (
	"github.com/DylanCoon99/server-with-rate-limiter/src/client"
)



func main() {

	// create an instance of the client and send requests from here

	client.CreateClient()

	client.SendRequest("http://localhost:8080/api/test", "GET", "")
}