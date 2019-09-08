package rubric_test

import (
	. "github.com/egoholic/blog/domain/rubric"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rubric", func() {
	Describe("accessors", func() {
		attrs := Attrs{
			Slug:            "slug",
			MetaKeywords:    "some, rubric",
			MetaDescription: "Some rubric",
			Title:           "Some Rubric",
			Description:     "Some rubric description.",
		}
		rubric := New(&attrs)

		Describe(".Attrs()", func() {
			It("returns meta-keywords", func() {
				Expect(rubric.Attrs()).To(Equal(&attrs))
			})
		})
	})
})
