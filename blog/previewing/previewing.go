package previewing

import (
	"log"

	"github.com/egoholic/blog/meta"
)

type (
	Blog struct {
		Title       string
		Keywords    string
		Description string
	}
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
		blog                        *Blog
		popularPublicationsProvider PopularPublicationsProvider
		recentPublicationsProvider  RecentPublicationsProvider
		rubricsProvider             RubricsProvider
		Meta                        *meta.Meta
	}
	BlogProvider interface {
		BlogByDomain(string) (*Blog, error)
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

func New(l *log.Logger, domain string, bp BlogProvider, ppp PopularPublicationsProvider, rpp RecentPublicationsProvider, rp RubricsProvider) (*Value, error) {
	blog, err := bp.BlogByDomain(domain)
	if err != nil {
		return nil, err
	}
	return &Value{
		logger:                      l,
		blog:                        blog,
		popularPublicationsProvider: ppp,
		recentPublicationsProvider:  rpp,
		rubricsProvider:             rp,
		Meta: &meta.Meta{
			Title:           blog.Title,
			MetaKeywords:    blog.Keywords,
			MetaDescription: blog.Description,
		},
	}, nil
}
func (v *Value) Blog() *Blog {
	return v.blog
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
