package reading

import "log"

type (
	Author struct {
		FullName string
		Bio      string
		Login    string
	}
	Publication struct {
		Slug       string
		Title      string
		Content    string
		CreatedAt  string
		Popularity int
	}
	Value struct {
		logger              *log.Logger
		publicationProvider PublicationProvider
		authorsProvider     AuthorsProvider
		slug                string
	}
	PublicationProvider interface {
		PublicationBySlug(string) (*Publication, error)
	}
	AuthorsProvider interface {
		AuthorsOf(string) ([]*Author, error)
	}
)

func New(l *log.Logger, pp PublicationProvider, ap AuthorsProvider, s string) *Value {
	return &Value{
		logger:              l,
		publicationProvider: pp,
		authorsProvider:     ap,
		slug:                s,
	}
}

func (v *Value) Publication() *Publication {
	publication, err := v.publicationProvider.PublicationBySlug(v.slug)
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return publication
}

func (v *Value) Authors() []*Author {
	authors, err := v.authorsProvider.AuthorsOf(v.slug)
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return authors
}
