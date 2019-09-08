package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/egoholic/blog/domain/publication"
	"github.com/egoholic/blog/domain/rubric"
	_ "github.com/lib/pq"
)

const RECENT_NUMBER = 10

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetRecentPublications(ctx context.Context) (publications []*publication.Publication) {
	rows, err := r.db.QueryContext(ctx, fmt.Sprintf("SELECT slug, meta_keywords, meta_description, title, content, created_at FROM publications ORDER BY created_at DESC LIMIT %d;", RECENT_NUMBER))
	if err != nil {
		panic(err)
	}
	if rows != nil {
		defer rows.Close()
	} else {
		panic(errors.New("no rows error"))
	}
	for rows.Next() {
		var attrs publication.Attrs
		err := rows.Scan(&attrs.Slug, &attrs.MetaKeywords, &attrs.MetaDescription, &attrs.Title, &attrs.Content, &attrs.CreatedAt)
		if err != nil {
			fmt.Printf("ERR: %#v", publications)
			panic(err)
		}
		publications = append(publications, publication.New(&attrs))
	}
	err = rows.Err()
	if err != nil {
		panic(err)
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
		panic(err)
	}
	return rubrics
}
