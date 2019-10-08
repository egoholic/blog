package previewing

import (
	"log"
)

type (
	Rubric struct {
		Slug  string
		Title string
	}
	Publication struct {
		Slug       string
		Title      string
		CreatedAt  string
		Popularity int
	}
	Value struct {
		logger                      *log.Logger
		popularPublicationsProvider PopularPublicationsProvider
		recentPublicationsProvider  RecentPublicationsProvider
		rubricsProvider             RubricsProvider
	}
	RecentPublicationsProvider interface {
		RecentPublications() ([]*Publication, error)
	}
	PopularPublicationsProvider interface {
		PopularPublications() ([]*Publication, error)
	}
	RubricsProvider interface {
		Rubrics() ([]*Rubric, error)
	}
)

func New(l *log.Logger, ppp PopularPublicationsProvider, rpp RecentPublicationsProvider, rp RubricsProvider) *Value {
	return &Value{
		logger:                      l,
		popularPublicationsProvider: ppp,
		recentPublicationsProvider:  rpp,
		rubricsProvider:             rp,
	}
}

func (v *Value) PopularPublications() []*Publication {
	publications, err := v.popularPublicationsProvider.PopularPublications()
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return publications
}

func (v *Value) RecentPublications() []*Publication {
	publications, err := v.recentPublicationsProvider.RecentPublications()
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return publications
}

func (v *Value) Rubrics() []*Rubric {
	rubrics, err := v.rubricsProvider.Rubrics()
	if err != nil {
		v.logger.Printf("ERROR: %s\n", err.Error())
	}
	return rubrics
}
