package handlers


import (
	"log"
	"fmt"
	"html"
	"net/http"
	"time"
)


// note: done channels are a channel used to notify another goroutine that the functions work is done
// 


func GetHandler(w http.ResponseWriter, r *http.Request) {


	ctx := r.Context()


	log.Printf("Hey I got the request! Here is the method: %v", r.Method)

	for i := 0; i < 10; i ++ {
		select {
		case <-ctx.Done(): // if done is closed --> context has ended  ; so if this channel receives it indicates the context is done
			log.Printf("Client ctx ended. Error: %v", ctx.Err())
			return
		default:
			log.Printf("Simulating work: %v s", i)
			time.Sleep(time.Second)
		}
	}
	log.Println("Work is done!")
	return
}


func Get2Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
