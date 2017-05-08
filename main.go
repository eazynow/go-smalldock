package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-zoo/bone"
)

const (
	// RFC3339Mili is a modification of RFC3339Nano to only include ms (3dp)
	RFC3339Mili = "2006-01-02T15:04:05.999Z07:00"
	// HTTPPort is the port to run the server on
	HTTPPort = 8080
)

// FormatTime formats a time object to RFC3339 with ms precision
func formatTime(timestamp time.Time) string {
	return timestamp.UTC().Format(RFC3339Mili)
}

func main() {
	mux := bone.New()
	mux.Get("/health", http.HandlerFunc(healthHandler))
	fmt.Printf("listening on localhost, port %d...", HTTPPort)
	http.ListenAndServe(fmt.Sprintf(":%d", HTTPPort), mux)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK at %s", time.Now())
}
