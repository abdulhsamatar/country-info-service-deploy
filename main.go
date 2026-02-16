package main

import (
	"log"
	"net/http"
	"time"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/handlers"
)

var startTime = time.Now()

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/countryinfo/v1/status/", handlers.StatusHandler(startTime))
	mux.HandleFunc("/countryinfo/v1/info/", handlers.InfoHandler)
	mux.HandleFunc("/countryinfo/v1/exchange/", handlers.ExchangeHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
