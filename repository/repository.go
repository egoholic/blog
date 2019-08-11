package repository

import (
	"context"

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

func (r *Repository) GetRecentPublications(ctx context.Context) (publications []*publication.Publication) {
	db := r.connector.Connection()
	defer db.Close()

	rows, err := db.Query(`SELECT meta_keywords, meta_description, title, content, created_at
						             FROM publications
						             ORDER created_at BY DESC
					               LIMIT 10;`)
	defer rows.Close()
	for rows.Next() {
		var attrs publication.Attrs
		err := rows.Scan(&attrs.MetaKeywords, &attrs.MetaDescription, &attrs.Title, &attrs.Content, &attrs.CreatedAt)
		if err != nil {

		}
		publications = append(publications, publication.New(&attrs))
	}
	err = rows.Err()
	if err != nil {

	}
	return publications
}

func (r *Repository) GetAllRubrics(ctx context.Context) (rubrics []*rubric.Rubric) {
	db := r.connector.Connection()
	defer db.Close()

	rows, err := db.Query(`SELECT meta_keywords, meta_description, title, description
						             FROM rubric
						             ORDER title BY ASC;`)
	defer rows.Close()
	for rows.Next() {
		var attrs rubric.Attrs
		err := rows.Scan(&attrs.MetaKeywords, &attrs.MetaDescription, &attrs.Title, &attrs.Description)
		if err != nil {

		}
		rubrics = append(rubrics, rubric.New(&attrs))
	}
	err = rows.Err()
	if err != nil {

	}
	return rubrics
}
