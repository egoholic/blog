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

var (
	blogQuery = `SELECT title
							 FROM (SELECT domain,
														title
										 FROM blogs
										 WHERE domain = $1
										 LIMIT 1) AS b;`
	rubricQuery = `SELECT slug,
											  	title,
													description,
													meta_keywords,
													meta_description
								   FROM rubrics
								   WHERE slug = $1
									 LIMIT 1;`
	publicationsQuery = `SELECT slug,
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
)

func (r *Repository) RubricBySlug(s string) (*previewing.Rubric, error) {
	var rubric previewing.Rubric
	row := r.db.QueryRowContext(r.ctx, rubricQuery, s)
	err := row.Scan(&rubric.Slug, &rubric.Title, &rubric.Description, &rubric.MetaKeywords, &rubric.MetaDescription)
	return &rubric, err
}
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
func (r *Repository) BlogByDomain(domain string) (*previewing.Blog, error) {
	var blog previewing.Blog
	row := r.db.QueryRowContext(r.ctx, blogQuery, domain)
	err := row.Scan(&blog.Title)
	return &blog, err
}
