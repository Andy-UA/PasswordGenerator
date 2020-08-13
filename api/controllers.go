package api

import (
	"fmt"
	"net/http"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func HandleRequests() {
	passwordConfig := http.HandlerFunc(getPasswordConfigs)

	//added middleware to validate request and use context for passing valid data
	http.Handle("/generate", httpMethodHandler(bodyAvailabilityHandler(bodyContentHandler(passwordConfig))))
	http.HandleFunc("/", sayHelloHandler)
}
