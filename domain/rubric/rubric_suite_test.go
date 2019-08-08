package rubric_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRubric(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rubric Suite")
}
