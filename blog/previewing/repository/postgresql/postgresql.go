package postgresql

import (
	"context"
	"database/sql"
	"log"

	"github.com/egoholic/blog/blog/previewing"
	_ "github.com/lib/pq"
)

type Repository struct {
	db     *sql.DB
	ctx    context.Context
	logger *log.Logger
}

var (
	blogQuery = `SELECT title,
											keywords,
											description
							 FROM (SELECT domain,
														title,
														keywords,
														description
										 FROM blogs
										 WHERE domain = $1
										 LIMIT 1) AS b;`
	recentPublicationsQuery = `SELECT slug,
	                                  title,
	                                  created_at,
																		popularity
														FROM publications ORDER BY created_at DESC LIMIT $1;`
	popularPublicationsQuery = `SELECT slug,
	                                   title,
	                                   created_at,
																		 popularity
															FROM publications ORDER BY popularity DESC LIMIT $1;`
	rubricsQuery = `SELECT slug, title FROM rubrics ORDER BY title ASC;`
)

func New(ctx context.Context, db *sql.DB, logger *log.Logger) *Repository {
	return &Repository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

func (r *Repository) RecentPublications() (publications []*previewing.Publication, err error) {
	rows, err := r.db.QueryContext(r.ctx, recentPublicationsQuery, 5)
	if err != nil {
		r.logger.Panicf("ERROR: %s\n", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var p previewing.Publication
		err = rows.Scan(&p.Slug, &p.Title, &p.CreatedAt, &p.Popularity)
		if err != nil {
			return []*previewing.Publication{}, err
		}
		publications = append(publications, &p)
	}
	return
}

func (r *Repository) PopularPublications() (publications []*previewing.Publication, err error) {
	rows, err := r.db.QueryContext(r.ctx, popularPublicationsQuery, 5)
	if err != nil {
		r.logger.Panicf("ERROR: %s\n", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var p previewing.Publication
		err = rows.Scan(&p.Slug, &p.Title, &p.CreatedAt, &p.Popularity)
		if err != nil {
			return []*previewing.Publication{}, err
		}
		publications = append(publications, &p)
	}
	return
}

func (r *Repository) Rubrics() (rubrics []*previewing.Rubric, err error) {
	rows, err := r.db.QueryContext(r.ctx, rubricsQuery)
	if err != nil {
		r.logger.Panicf("ERROR: %s\n", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var r previewing.Rubric
		err = rows.Scan(&r.Slug, &r.Title)
		if err != nil {
			return []*previewing.Rubric{}, err
		}
		rubrics = append(rubrics, &r)
	}
	return
}

func (r *Repository) BlogByDomain(domain string) (*previewing.Blog, error) {
	var b previewing.Blog
	row := r.db.QueryRowContext(r.ctx, blogQuery, domain)
	err := row.Scan(&b.Title, &b.Keywords, &b.Description)
	return &b, err
}
