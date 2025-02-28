package services

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOutboxServiceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain suite")
}
