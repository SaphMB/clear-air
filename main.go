package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SaphMB/clear-air/airqualityclient"
)

func main() {
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(":8080", nil))

	apiToken := os.Getenv("WORLD_AIR_QUALITY_API_TOKEN")
	client := &airqualityclient.AirQualityClient{APIToken: apiToken}
	client.Validate()
}

func handleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Listening on %s 🎶", req.Host)
}
