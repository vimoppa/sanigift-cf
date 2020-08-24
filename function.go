package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

var d struct {
	Message string `json:"message"`
}

// Uploader func dumps file into Storage bucket
func Uploader(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprintf(w, "Hello World")
		return
	}
	if d.Message == "" {
		fmt.Fprintf(w, "Hello, World!")
		return
	}
	fmt.Fprintf(w, html.EscapeString(d.Message))
}
