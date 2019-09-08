package publication_test

import (
	. "github.com/egoholic/blog/domain/publication"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("publication", func() {
	Describe("accessors", func() {
		attrs := Attrs{
			Slug:            "slug",
			MetaKeywords:    "some, publication",
			MetaDescription: "Some publication",
			Title:           "Some Publication",
			Content:         "Some publication content.",
			CreatedAt:       "",
		}
		puplication := New(&attrs)
		Describe(".Attrs()", func() {
			It("returns meta-keywords", func() {
				Expect(puplication.Attrs()).To(Equal(&attrs))
			})
		})
	})
})
