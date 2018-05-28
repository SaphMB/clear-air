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
			status, err := client.Validate()
			Expect(err).ToNot(HaveOccurred())
			Expect(status).To(Equal("ok"))
		})
	})
})
