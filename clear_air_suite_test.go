package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClearAir(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ClearAir Suite")
}
