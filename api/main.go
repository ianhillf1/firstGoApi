package main

import (
    "log"
    "net/http"
	"./routes"
)

func main() {
	router:= routes.BuildRouter()
	log.Fatal(http.ListenAndServe(":10000", router))
}