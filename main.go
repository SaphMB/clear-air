package main

import (
	"os"

	"github.com/SaphMB/clear-air/airqualityclient"
	"github.com/SaphMB/clear-air/api"
)

func main() {

	go api.Start()

	apiToken := os.Getenv("WORLD_AIR_QUALITY_API_TOKEN")
	client := &airqualityclient.AirQualityClient{APIToken: apiToken}
	client.Validate()
}
