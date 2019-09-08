package repository_test

import (
	"context"
	"database/sql"
	"fmt"

	. "github.com/egoholic/blog/config"
	. "github.com/egoholic/blog/repository"
	. "github.com/egoholic/blog/store/seed"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("repository", func() {
	connStr, _ := Config.DBCredentials().ConnectionString()
	db, _ := sql.Open("postgres", connStr)
	//defer db.Close()
	repo := New(db)
	Describe("New()", func() {
		BeforeSuite(func() {
			Truncate(db, "publications", "rubrics")
		})
		It("returns repository", func() {
			Expect(repo).NotTo(BeNil())
			Expect(repo).To(BeAssignableToTypeOf(&Repository{}))
		})
	})
	Describe("Repository", func() {
		Describe(".GetRecentPublications()", func() {
			Context(fmt.Sprintf("when there are more than %d publications", RECENT_NUMBER), func() {
				It(fmt.Sprintf("returns only %d publications", RECENT_NUMBER), func() {
					_, err := Many(RECENT_NUMBER+1, db, CreatePublication)
					if err != nil {
						panic(err)
					}
					publications := repo.GetRecentPublications(context.Background())
					Expect(publications).To(HaveLen(10))
				})
			})
		})
		Describe(".GetAllRubrics()", func() {
		})
	})
})
