package previewing

import (
	"fmt"
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
		Slug       string
		Title      string
		CreatedAt  string
		Popularity int
	}
	Value struct {
		logger               *log.Logger
		author               *Author
		login                string
		Meta                 *meta.Meta
		publicationsProvider PublicationsProvider
	}
	AuthorProvider interface {
		AuthorByLogin(l string) (*Author, error)
	}
	PublicationsProvider interface {
		PublicationsOf(l string) ([]*Publication, error)
	}
)

func New(l *log.Logger, ap AuthorProvider, pp PublicationsProvider, login string) (*Value, error) {
	author, err := ap.AuthorByLogin(login)
	title := fmt.Sprintf("%s author's page", author.FullName)
	if err != nil {
		return nil, err
	}
	return &Value{
		logger: l,
		author: author,
		login:  login,
		Meta: &meta.Meta{
			Title:           title,
			MetaKeywords:    "author, publications",
			MetaDescription: title,
		},
		publicationsProvider: pp,
	}, nil
}
func (v *Value) Publications() []*Publication {
	publications, err := v.publicationsProvider.PublicationsOf(v.login)
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return publications
}
func (v *Value) Author() *Author {
	return v.author
}
