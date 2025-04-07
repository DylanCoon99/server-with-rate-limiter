package middleware



import (
	//"fmt"
	"log"
	"time"
	"net/http"
)


type Logger struct {
	handler http.Handler
}


func (logger *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	logger.handler.ServeHTTP(w, r)
	log.Printf("%v %v %v", r.Method, r.URL.Path, time.Since(start))

}



func NewLogger(handlerToWrap http.Handler) *Logger {
	// This function wraps a handler in the logger
	return &Logger{handlerToWrap}
}




