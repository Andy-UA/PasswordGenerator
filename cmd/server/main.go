package main

import (
	"../../api"
	"log"
	"net/http"
)

func main() {
	api.HandleRequests()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
