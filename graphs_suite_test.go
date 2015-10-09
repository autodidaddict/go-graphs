package graphs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCfenv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Graphs Test Suite")
}
