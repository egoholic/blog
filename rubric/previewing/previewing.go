package previewing

import "log"

type (
	Rubric struct {
		Slug        string
		Title       string
		Description string
	}
	Publication struct {
		Slug       string
		Title      string
		CreatedAt  string
		Popularity string
	}
	Value struct {
		logger               *log.Logger
		rubricProvider       RubricProvider
		publicationsProvider PublicationsProvider
		slug                 string
	}
	RubricProvider interface {
		RubricBySlug(string) (*Rubric, error)
	}
	PublicationsProvider interface {
		PublicationsOf(string) ([]*Publication, error)
	}
)

func New(l *log.Logger, rp RubricProvider, pp PublicationsProvider, slug string) *Value {
	return &Value{
		logger:               l,
		rubricProvider:       rp,
		publicationsProvider: pp,
		slug:                 slug,
	}
}

func (v *Value) Rubric() *Rubric {
	rubric, err := v.rubricProvider.RubricBySlug(v.slug)
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return rubric
}

func (v *Value) Publications() []*Publication {
	publications, err := v.publicationsProvider.PublicationsOf(v.slug)
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return publications
}
