package publication_test

import (
	. "github.com/egoholic/blog/domain/publication"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("publication", func() {
	Describe("accessors", func() {
		var (
			meta        = Meta{Keywords: []string{"some", "publication"}, Description: "Some publication"}
			attrs       = Attrs{Meta: meta, Title: "Some Publication", Content: "Some publication content."}
			puplication = New(&attrs)
		)
		Describe(".MetaKeywords()", func() {
			It("returns meta-keywords", func() {
				Expect(puplication.MetaKeywords()).To(Equal([]string{"some", "publication"}))
			})
		})
		Describe(".MetaDescription()", func() {
			It("returns meta-description", func() {
				Expect(puplication.MetaDescription()).To(Equal("Some publication"))
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				Expect(puplication.Title()).To(Equal("Some Publication"))
			})
		})
		Describe(".Content()", func() {
			It("returns content", func() {
				Expect(puplication.Content()).To(Equal("Some publication content."))
			})
		})
	})
})
