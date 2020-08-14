package api

import (
	"net/http"
)

func HandleRequests() {
	passwordConfig := http.HandlerFunc(getPasswordConfigs)

	//added middleware to validate request and use context for passing valid data
	http.Handle("/generate", httpMethodHandler(bodyAvailabilityHandler(bodyContentHandler(passwordConfig))))
}
