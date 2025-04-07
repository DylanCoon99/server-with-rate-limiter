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

	* Let's start with simply logging all requests (Done)
	* I will implement the api limiter using the fixed window algorithm
*/


func CreateServer(port string) {

	// here we are going to instantiate a server 

	serverMux := http.NewServeMux()

	// Put handle funcs here
	serverMux.HandleFunc("GET /api/test", handlers.GetHandler2)

	SERVER = &http.Server{
		Addr: ":" + port,
		Handler: serverMux,
	}


	wrappedMux := middleware.NewCustomHandler(serverMux)

	Listen("localhost:" + port, wrappedMux)

}



func Listen(port string, mux *middleware.CustomHandler) {

	log.Printf("Listening on port: %v", SERVER.Addr)
	log.Fatal(http.ListenAndServe(port, mux))

}