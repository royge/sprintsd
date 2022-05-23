package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/royge/sprintsd"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := "sprintsd"
	uri, err := sprintsd.GetRunServiceURL(context.Background(), name)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Service URL: %s!\n", uri)
}
