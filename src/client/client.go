package client

import (
	"fmt"
	"time"
	"context"
	"net/http"
)


var CLIENT *http.Client


func CreateClient() {


	CLIENT = &http.Client{
		CheckRedirect: nil,
	}

	fmt.Println("Client has been instantiated")
}


func SendRequest(url string, method string, payload string) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancelFunc()

	switch method {
	case "GET":
		// handle GET request
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			fmt.Printf("Failed to create GET request: %v\n", err)
			return
		}

		resp, err := CLIENT.Do(req)
		if err != nil {
			fmt.Printf("Failed to carry out GET request: %v", err)
			return
		}

		fmt.Println(resp.StatusCode)
	case "POST":
		// handle POST request
		// To be implemented
	default:
		fmt.Println("Sorry, you can only send GET and POST requests.")
		return
	}


	return

}