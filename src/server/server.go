package server

import (
	//"fmt"
	"log"
	"net/http"
	"github.com/DylanCoon99/server-with-rate-limiter/src/handlers"
)


var SERVER *http.Server

// want the server to be aware of the context of the request
// - whether or not the client is connected or not
// - API rate limiting


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