package previewing

import "log"

type (
	Author struct {
		Login    string
		FullName string
		Bio      string
	}
	Publication struct {
		Slug       string
		Title      string
		CreatedAt  string
		Popularity int
	}
	Value struct {
		logger               *log.Logger
		authorProvider       AuthorProvider
		publicationsProvider PublicationsProvider
		login                string
	}
	AuthorProvider interface {
		AuthorByLogin(l string) (*Author, error)
	}
	PublicationsProvider interface {
		PublicationsOf(l string) ([]*Publication, error)
	}
)

func New(l *log.Logger, ap AuthorProvider, pp PublicationsProvider, login string) *Value {
	return &Value{
		logger:               l,
		authorProvider:       ap,
		publicationsProvider: pp,
		login:                login,
	}
}

func (v *Value) Publications() []*Publication {
	publications, err := v.publicationsProvider.PublicationsOf(v.login)
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return publications
}

func (v *Value) Author() *Author {
	author, err := v.authorProvider.AuthorByLogin(v.login)
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return author
}
