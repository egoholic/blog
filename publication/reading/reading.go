package reading

import (
	"log"

	"github.com/egoholic/blog/meta"
)

type (
	Blog struct {
		Title string
	}
	Author struct {
		Login    string
		FullName string
		Bio      string
	}
	Publication struct {
		Slug            string
		Title           string
		Content         string
		CreatedAt       string
		Popularity      int
		MetaKeywords    string
		MetaDescription string
	}
	Value struct {
		logger          *log.Logger
		publication     *Publication
		blog            *Blog
		authorsProvider AuthorsProvider
		slug            string
		Meta            *meta.Meta
	}
	PublicationProvider interface {
		PublicationBySlug(string) (*Publication, error)
	}
	BlogProvider interface {
		BlogByDomain(string) (*Blog, error)
	}
	AuthorsProvider interface {
		AuthorsOf(string) ([]*Author, error)
	}
)

func New(l *log.Logger, pp PublicationProvider, bp BlogProvider, ap AuthorsProvider, slug, domain string) (*Value, error) {
	publication, err := pp.PublicationBySlug(slug)
	if err != nil {
		return nil, err
	}
	blog, err := bp.BlogByDomain(domain)
	if err != nil {
		return nil, err
	}
	return &Value{
		logger:          l,
		publication:     publication,
		blog:            blog,
		authorsProvider: ap,
		slug:            slug,
		Meta: &meta.Meta{
			Title:           publication.Title,
			MetaDescription: publication.MetaDescription,
			MetaKeywords:    publication.MetaKeywords,
		},
	}, nil
}
func (v *Value) Publication() *Publication {
	return v.publication
}
func (v *Value) Authors() []*Author {
	authors, err := v.authorsProvider.AuthorsOf(v.slug)
	if err != nil {
		v.logger.Printf("ERROR-publication-reading: %s\n", err.Error())
	}
	return authors
}
func (v *Value) Blog() *Blog {
	return v.blog
}
