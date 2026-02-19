package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/handlers"
)

var startTime time.Time

func main() {
	// Record service start time
	startTime = time.Now()

	// Set up routes
	http.HandleFunc("/countryinfo/v1/status/", handlers.StatusHandler(startTime))
	http.HandleFunc("/countryinfo/v1/info/", handlers.InfoHandler)
	http.HandleFunc("/countryinfo/v1/exchange/", handlers.ExchangeHandler)

	// Root handler for documentation/help
	http.HandleFunc("/", rootHandler)

	// Start server
	port := "8080"
	fmt.Printf("Starting Country Info Service on port %s...\n", port)
	fmt.Printf("Endpoints:\n")
	fmt.Printf("  - GET /countryinfo/v1/status/\n")
	fmt.Printf("  - GET /countryinfo/v1/info/{code}\n")
	fmt.Printf("  - GET /countryinfo/v1/exchange/{code}\n")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// rootHandler provides basic API documentation
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Only handle exact root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	help := `Country Info Service API v1

Available endpoints:

GET /countryinfo/v1/status/
GET /countryinfo/v1/info/{code}
GET /countryinfo/v1/exchange/{code}

Example:
curl http://localhost:8080/countryinfo/v1/info/no
`

	fmt.Fprint(w, help)
}
