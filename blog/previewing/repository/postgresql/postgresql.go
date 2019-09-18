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
		r.logger.Println("recent-fuck")
		return
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
		r.logger.Println("popular-fuck")

		return
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
