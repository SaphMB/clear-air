package api

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Listening on %s 🎶", req.Host)
}
