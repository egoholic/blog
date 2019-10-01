package postgresql

import (
	"context"
	"database/sql"
	"log"

	"github.com/egoholic/blog/rubric/previewing"
	_ "github.com/lib/pq"
)

type Repository struct {
	db     *sql.DB
	ctx    context.Context
	logger *log.Logger
}

func New(ctx context.Context, db *sql.DB, logger *log.Logger) *Repository {
	return &Repository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

var rubricQuery = `SELECT slug,
											  	title,
												  description
								   FROM rubrics
								   WHERE slug = $1
								   LIMIT 1;`

func (r *Repository) RubricBySlug(s string) (*previewing.Rubric, error) {
	var rubric previewing.Rubric
	row := r.db.QueryRowContext(r.ctx, rubricQuery, s)
	err := row.Scan(&rubric.Slug, &rubric.Title, &rubric.Description)
	return &rubric, err
}

var publicationsQuery = `SELECT slug,
															  title,
															  created_at,
															  popularity
											   FROM (SELECT slug,
															        title,
														  	      created_at,
															        popularity,
															        rubric_slug
											         FROM publications
															 WHERE rubric_slug = $1) AS selected
												 ORDER BY created_at DESC;`

func (r *Repository) PublicationsOf(s string) (publications []*previewing.Publication, err error) {
	rows, err := r.db.QueryContext(r.ctx, publicationsQuery, s)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p previewing.Publication
		err = rows.Scan(&p.Slug, &p.Title, &p.CreatedAt, &p.Popularity)
		if err != nil {
			return
		}
		publications = append(publications, &p)
	}
	return
}
