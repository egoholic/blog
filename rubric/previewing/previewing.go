package previewing

import (
	"log"

	"github.com/egoholic/blog/meta"
)

type (
	Rubric struct {
		Slug            string
		Title           string
		Description     string
		MetaKeywords    string
		MetaDescription string
	}
	Publication struct {
		Slug       string
		Title      string
		CreatedAt  string
		Popularity string
	}
	Value struct {
		logger               *log.Logger
		rubric               *Rubric
		publicationsProvider PublicationsProvider
		slug                 string
		Meta                 *meta.Meta
	}
	RubricProvider interface {
		RubricBySlug(string) (*Rubric, error)
	}
	PublicationsProvider interface {
		PublicationsOf(string) ([]*Publication, error)
	}
)

func New(l *log.Logger, rp RubricProvider, pp PublicationsProvider, slug string) (*Value, error) {
	rubric, err := rp.RubricBySlug(slug)
	if err != nil {
		return nil, err
	}
	return &Value{
		logger:               l,
		rubric:               rubric,
		publicationsProvider: pp,
		slug:                 slug,
		Meta: &meta.Meta{
			Title:           rubric.Title,
			MetaKeywords:    rubric.MetaKeywords,
			MetaDescription: rubric.MetaDescription,
		},
	}, nil
}

func (v *Value) Rubric() *Rubric {
	return v.rubric
}

func (v *Value) Publications() []*Publication {
	publications, err := v.publicationsProvider.PublicationsOf(v.slug)
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return publications
}
