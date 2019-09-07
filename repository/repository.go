package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/egoholic/blog/domain/publication"
	"github.com/egoholic/blog/domain/rubric"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetRecentPublications(ctx context.Context) (publications []*publication.Publication) {
	rows, err := r.db.QueryContext(ctx, `SELECT meta_keywords, meta_description, title, content, created_at
						             FROM publications
						             ORDER created_at BY DESC
					               LIMIT 10;`)
	defer rows.Close()
	for rows.Next() {
		var attrs publication.Attrs
		var keywords string
		err := rows.Scan(&keywords, &attrs.MetaDescription, &attrs.Title, &attrs.Content, &attrs.CreatedAt)
		attrs.MetaKeywords = strings.Split(keywords, ",")
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
	rows, err := r.db.QueryContext(ctx, `SELECT meta_keywords, meta_description, title, description
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
