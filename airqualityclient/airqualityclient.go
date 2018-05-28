package airqualityclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AirQualityClient struct {
	APIToken string
}

type CityFeed struct {
	Status  string `json:"status"`
	Data    CityData
	Message string
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

func (a AirQualityClient) Validate() (string, error) {
	token := a.APIToken
	resp, err := http.Get("http://api.waqi.info/feed/@5724/?token=" + token)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	status, statusErr := a.statusCheck(resp)

	return status, statusErr
}

func (a AirQualityClient) statusCheck(responseBody *http.Response) (string, error) {

	body, readErr := ioutil.ReadAll(responseBody.Body)
	if readErr != nil {
		log.Fatalf("Error: %s", readErr)
	}
	defer responseBody.Body.Close()

	feed := &CityFeed{}
	unmarshallErr := json.Unmarshal(body, feed)
	if unmarshallErr != nil {
		log.Fatalf("Error unmarshalling response: %s", unmarshallErr)
	}
	status := feed.Status
	if status == "error" {
		log.Fatalf("Error: %s", feed.Message)
	}

	return status, unmarshallErr
}
