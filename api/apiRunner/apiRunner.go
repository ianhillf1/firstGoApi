package main

import (
    "log"
    "net/http"
	".."
)

func main() {
	router:= api.BuildRouter()
	log.Fatal(http.ListenAndServe(":10000", router))
}