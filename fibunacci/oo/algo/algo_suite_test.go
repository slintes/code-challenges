package algo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAlgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Algo Suite")
}
