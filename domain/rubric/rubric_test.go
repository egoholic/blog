package rubric_test

import (
	. "github.com/egoholic/blog/domain/rubric"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rubric", func() {
	Describe("accessors", func() {
		var (
			meta   = Meta{Keywords: []string{"some", "rubric"}, Description: "Some rubric"}
			attrs  = Attrs{Meta: meta, Title: "Some Rubric", Content: "Some rubric content."}
			rubric = New(&attrs)
		)
		Describe(".MetaKeywords()", func() {
			It("returns meta-keywords", func() {
				Expect(rubric.MetaKeywords()).To(Equal([]string{"some", "rubric"}))
			})
		})
		Describe(".MetaDescription()", func() {
			It("returns meta-description", func() {
				Expect(rubric.MetaDescription()).To(Equal("Some rubric"))
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				Expect(rubric.Title()).To(Equal("Some Rubric"))
			})
		})
		Describe(".Content()", func() {
			It("returns content", func() {
				Expect(rubric.Content()).To(Equal("Some rubric content."))
			})
		})
	})
})
