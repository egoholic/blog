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
	BeforeEach(func() {
		Truncate(db, "publications", "rubrics")
	})
	AfterEach(func() {
		Truncate(db, "publications", "rubrics")
	})
	Describe("New()", func() {
		It("returns repository", func() {
			Expect(repo).NotTo(BeNil())
			Expect(repo).To(BeAssignableToTypeOf(&Repository{}))
		})
	})
	Describe("Repository", func() {
		Describe(".GetRecentPublications()", func() {
			Context(fmt.Sprintf("when there are more than %d publications", RECENT_NUMBER), func() {
				It(fmt.Sprintf("returns only %d publications", RECENT_NUMBER), func() {
					Many(RECENT_NUMBER+1, db, CreatePublication)
					publications := repo.GetRecentPublications(context.Background())
					Expect(publications).To(HaveLen(10))
					Expect(publications[0].Attrs().Title).To(Equal("10th PUBLICATION"))
					Expect(publications[1].Attrs().Title).To(Equal("9th PUBLICATION"))
					Expect(publications[2].Attrs().Title).To(Equal("8th PUBLICATION"))
					Expect(publications[3].Attrs().Title).To(Equal("7th PUBLICATION"))
					Expect(publications[4].Attrs().Title).To(Equal("6th PUBLICATION"))
					Expect(publications[5].Attrs().Title).To(Equal("5th PUBLICATION"))
					Expect(publications[6].Attrs().Title).To(Equal("4th PUBLICATION"))
					Expect(publications[7].Attrs().Title).To(Equal("3th PUBLICATION"))
					Expect(publications[8].Attrs().Title).To(Equal("2th PUBLICATION"))
					Expect(publications[9].Attrs().Title).To(Equal("1th PUBLICATION"))
				})
			})

			Context(fmt.Sprintf("when there are less than %d publications", RECENT_NUMBER), func() {
				It(fmt.Sprintf("returns less than %d publications", RECENT_NUMBER), func() {
					Many(RECENT_NUMBER-1, db, CreatePublication)
					publications := repo.GetRecentPublications(context.Background())
					Expect(publications).To(HaveLen(9))
					Expect(publications[0].Attrs().Title).To(Equal("8th PUBLICATION"))
					Expect(publications[1].Attrs().Title).To(Equal("7th PUBLICATION"))
					Expect(publications[2].Attrs().Title).To(Equal("6th PUBLICATION"))
					Expect(publications[3].Attrs().Title).To(Equal("5th PUBLICATION"))
					Expect(publications[4].Attrs().Title).To(Equal("4th PUBLICATION"))
					Expect(publications[5].Attrs().Title).To(Equal("3th PUBLICATION"))
					Expect(publications[6].Attrs().Title).To(Equal("2th PUBLICATION"))
					Expect(publications[7].Attrs().Title).To(Equal("1th PUBLICATION"))
					Expect(publications[8].Attrs().Title).To(Equal("0th PUBLICATION"))
				})
			})
		})
		Describe(".GetAllRubrics()", func() {
		})
	})
})
