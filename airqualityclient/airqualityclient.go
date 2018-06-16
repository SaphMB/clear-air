package airqualityclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AirQualityClient struct {
	APIToken string
}

type Status struct {
	Status string `json:"status"`
}

type CityFeed struct {
	Status string
	Data   interface{}
}

type CityData struct {
	Idx int `json:"idx"`
	Aqi int `json:"aqi"`
	City
	Iaqi
}

type City struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Iaqi struct {
	Pm25 struct {
		V int `json:"v"`
	}
}

func (a AirQualityClient) Validate() string {
	token := a.APIToken
	resp, err := http.Get("http://api.waqi.info/feed/@5724/?token=" + token)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	status := a.statusCheck(resp)

	return status
}

func (a AirQualityClient) statusCheck(responseBody *http.Response) string {

	body, readErr := ioutil.ReadAll(responseBody.Body)
	if readErr != nil {
		log.Fatalf("Error: %s", readErr)
	}
	defer responseBody.Body.Close()

	cityFeed := &CityFeed{}
	unmarshallErr := json.Unmarshal(body, cityFeed)
	if unmarshallErr != nil {
		log.Fatalf("Error unmarshalling response: %s", unmarshallErr)
	}

	var status string

	switch cityFeed.Status {
	case "error":
		errorMessage := fmt.Sprintf("Error: %s", cityFeed.Data)
		status = errorMessage
	case "ok":
		status = "ok"
	}

	return status
}
