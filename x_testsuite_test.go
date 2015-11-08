package gotility_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGotility(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gotility Test Suite")
}
