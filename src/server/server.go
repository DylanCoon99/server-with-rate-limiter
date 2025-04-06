package server

import (
	//"fmt"
	"log"
	"net/http"
	"github.com/DylanCoon99/server-with-rate-limiter/src/handlers"
	"github.com/DylanCoon99/server-with-rate-limiter/src/middleware"

)


var SERVER *http.Server


// API rate limiting
/*

Rate limiting Algorithms
	- Fixed Window: Limits the number of requests within a fixed time period.
	- Sliding Window: Tracks requests over a rolling period of time, allowing more flexible
	handling of traffic spikes.
	- Leaky Bucket: Processes requests at a consistent, controlled rate, preventing sudden
	traffic surges.
	- Token Bucket: Allows bursts of requests as long as there are tokens available, offering
	more flexibility than the leaky bucket.

*/


func CreateServer(port string) {

	// here we are going to instantiate a server 

	serverMux := http.NewServeMux()

	// Put handle funcs here
	serverMux.HandleFunc("GET /api/test", handlers.GetHandler)

	SERVER = &http.Server{
		Addr: ":" + port,
		Handler: serverMux,
	}

}



func Listen() {

	log.Printf("Listening on port: %v", SERVER.Addr)
	log.Fatal(SERVER.ListenAndServe())

}