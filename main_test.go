package main_test

import (
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Clear air", func() {

	var cliPath string

	BeforeSuite(func() {
		var err error
		cliPath, err = Build("github.com/SaphMB/clear-air")
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("/", func() {
		It("should return a 200 status code", func() {
			command := exec.Command(cliPath)
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			resp, err := http.Get("http://localhost:8080")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			session.Kill()
		})

	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})

})
