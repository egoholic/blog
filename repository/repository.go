package repository

import (
	"github.com/egoholic/blog/domain/publication"
	"github.com/egoholic/blog/domain/rubric"
	"github.com/egoholic/blog/store/connector"
)

type Repository struct {
	connector *connector.Connector
}

func New(connector *connector.Connector) *Repository {
	return &Repository{connector: connector}
}
func (r *Repository) GetRecentPublications() (publications []*publication.Publication) {
}
func (r *Repository) GetAllRubrics() []*rubric.Rubric {
}
