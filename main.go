package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eazynow/bone"
	"github.com/urfave/negroni"
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
	// bone is a lightweight HTTP multiplexer for go
	mux := bone.New()
	mux.Get("/health", http.HandlerFunc(healthHandler))

	// use negroni to make it easy to add middleware in later
	// at low cost. Also provides basic logging out the box
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(fmt.Sprintf(":%d", HTTPPort))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// just return a 200 with the current time
	fmt.Fprintf(w, "OK at %s", time.Now())
}
