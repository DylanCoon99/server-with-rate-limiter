package main

// entry point for the client

import (
	"log"
	"time"
	"github.com/DylanCoon99/server-with-rate-limiter/src/client"
)



func main() {

	// create an instance of the client and send requests from here

	client.CreateClient()

	for i := 0; i < 20; i++ {
		client.SendRequest("http://localhost:8080/api/test", "GET", "")
	}


	for i := 0; i < 15; i ++ {
		log.Printf("Waiting %d s", i)
		time.Sleep(time.Second * 1)
	}

	client.SendRequest("http://localhost:8080/api/test", "GET", "")
}