package middleware



import (
	//"fmt"
	"log"
	"time"
	"net/http"
)



/*
How middleware works: Basically, we can create a custom http.Handler. The only rule for a handler is 
that it must implement "ServeHTTP(w http.ResponseWriter, r *http.Request)". For example we can create
a logger type that implements this method and we can wrap the regular handler with the new handler. 
These custom handlers can be chained together to do different things
*/


const (
	MAX_COUNT     = 10    // maximum number of requests per window
	WINDOW_LENGTH = 10.0  // window length in seconds
)



type CustomHandler struct {
	handler http.Handler
	maxCount int
	windowLength float32
	windowStartTime time.Time
	requestsSinceWindow int
}


func (myHandler *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	start := time.Now()


	if time.Since(myHandler.windowStartTime) >= WINDOW_LENGTH * time.Second {
		// reset the window
		myHandler.windowStartTime = time.Now()
		myHandler.requestsSinceWindow = 0
	}

	
	if myHandler.requestsSinceWindow < myHandler.maxCount {
		// increment the requestsSinceWindow
		myHandler.requestsSinceWindow += 1

		// fulfill the request
		myHandler.handler.ServeHTTP(w, r)
		log.Printf("%v %v %v", r.Method, r.URL.Path, time.Since(start))
		return
	}

	log.Printf("Request denied (rate limit exceeded): %v %v %v", r.Method, r.URL.Path, time.Since(start))

	http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)


	return

}



func NewCustomHandler(handlerToWrap http.Handler) *CustomHandler {
	// This function wraps a handler in the logger
	return &CustomHandler{
		handler: handlerToWrap,
		maxCount: MAX_COUNT,
		windowLength: WINDOW_LENGTH,
		windowStartTime: time.Now(),
	}
}