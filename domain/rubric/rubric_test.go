package rubric_test

import (
	. "github.com/egoholic/blog/domain/rubric"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rubric", func() {
	Describe("accessors", func() {
		var (
			attrs = Attrs{
				MetaKeywords:    []string{"some", "rubric"},
				MetaDescription: "Some rubric",
				Title:           "Some Rubric",
				Description:     "Some rubric description.",
			}
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
		Describe(".Description()", func() {
			It("returns description", func() {
				Expect(rubric.Description()).To(Equal("Some rubric description."))
			})
		})
	})
})
