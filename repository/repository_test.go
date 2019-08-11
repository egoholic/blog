package repository_test

import (
	"github.com/egoholic/blog/config"
	. "github.com/egoholic/blog/repository"
	"github.com/egoholic/blog/store/connector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("repository", func() {
	Describe("New()", func() {
		It("returns repository", func() {
			credsKeeper := &config.DBCredentials{Host: "localhost", Port: 5432, User: "postgres", Password: "", DBName: ""}
			connector := connector.New(credsKeeper)
			repo := New(connector)
			Expect(repo).NotTo(BeNil())
			Expect(repo).To(BeAssignableToTypeOf(&Repository{}))
		})
	})

	Describe("Repository", func() {
		Describe(".GetRecentPublications()", func() {

		})

		Describe(".GetAllRubrics()", func() {

		})
	})
})
