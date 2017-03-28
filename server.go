package main

import (
	"net/http"
	"os"

	"fmt"

	"encoding/json"

	"github.com/russross/blackfriday"
)

func main() {
	// Routes
	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("./public")))

	// Get the port on wich to run
	port := os.Getenv("PORT")
	// Set default port to 8000
	if port == "" {
		port = "8000"
	}

	fmt.Println("Running on port " + port)
	http.ListenAndServe(":"+port, nil)
}

// Functions that reads markdown and convert and send it
// as html to the client
func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	// Convert markdown to html
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))

	// Format the content to json
	js, _ := json.Marshal(string(markdown))

	// Set response as json
	rw.Header().Set("Content-Type", "application/json")

	rw.WriteHeader(200)
	rw.Write(js)
}
