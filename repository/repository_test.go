package repository_test

import (
	"database/sql"

	"github.com/egoholic/blog/config"
	. "github.com/egoholic/blog/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("repository", func() {
	Describe("New()", func() {
		It("returns repository", func() {
			creds := &config.DBCredentials{Host: "localhost", Port: 5432, User: "postgres", Password: "", DBName: ""}
			connStr, _ := creds.ConnectionString()
			db, _ := sql.Open("postgres", connStr)
			repo := New(db)
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
