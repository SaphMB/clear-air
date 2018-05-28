package airqualityclient_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAirQualityClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AirQualityClient Suite")
}
