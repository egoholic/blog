package reading

import (
	"log"

	"github.com/egoholic/blog/meta"
)

type (
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
		authorsProvider AuthorsProvider
		slug            string
		Meta            *meta.Meta
	}
	PublicationProvider interface {
		PublicationBySlug(string) (*Publication, error)
	}
	AuthorsProvider interface {
		AuthorsOf(string) ([]*Author, error)
	}
)

func New(l *log.Logger, pp PublicationProvider, ap AuthorsProvider, s string) (*Value, error) {
	publication, err := pp.PublicationBySlug(s)
	if err != nil {
		return nil, err
	}
	return &Value{
		logger:          l,
		publication:     publication,
		authorsProvider: ap,
		slug:            s,
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
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return authors
}
