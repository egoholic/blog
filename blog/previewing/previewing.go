package previewing

import "log"

type (
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
	}
	RecentPublicationsProvider interface {
		RecentPublications() ([]*Publication, error)
	}
	PopularPublicationsProvider interface {
		PopularPublications() ([]*Publication, error)
	}
)

func New(l *log.Logger, ppp PopularPublicationsProvider, rpp RecentPublicationsProvider) *Value {
	return &Value{
		logger: l,
		popularPublicationsProvider: ppp,
		recentPublicationsProvider:  rpp,
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
