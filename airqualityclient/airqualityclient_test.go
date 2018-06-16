package airqualityclient_test

import (
	"os"

	"github.com/SaphMB/clear-air/airqualityclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Air Quality Client", func() {

	Context("When correct api token is provided", func() {

		apiToken := os.Getenv("WORLD_AIR_QUALITY_API_TOKEN")
		client := &airqualityclient.AirQualityClient{APIToken: apiToken}

		It("validates the client", func() {
			status := client.Validate()
			Expect(status).To(Equal("ok"))
		})
	})

	Context("When an incorrect api token is provided", func() {

		apiToken := "incorrecttoken"
		client := &airqualityclient.AirQualityClient{APIToken: apiToken}

		It("raises an error stating that the incorrect token hase been used", func() {
			status := client.Validate()
			Expect(status).To(Equal("Error: Invalid key"))
		})
	})
})
