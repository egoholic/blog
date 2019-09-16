package postgresql

import (
	"context"
	"database/sql"
	"log"

	"github.com/egoholic/blog/blog/previewing"
	_ "github.com/lib/pq"
)

type Provider struct {
	db  *sql.DB
	ctx context.Context
	log log.Logger
}

var (
	recentPublicationsQuery = `SELECT slug,
	                                  title,
	                                  content,
	                                  created_at,
																		popularity
														FROM publications ORDER BY created_at DESC LIMIT $1;`
	popularPublicationsQuery = `SELECT slug,
	                                   title,
	                                   content,
	                                   created_at,
																		 popularity
															FROM publications ORDER BY popularity DESC LIMIT $1;`
)

func (p *Provider) RecentPublicationsProvider() (publications []*previewing.Publication, err error) {
	rows, err := p.db.QueryContext(p.ctx, recentPublicationsQuery, 5)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p previewing.Publication
		err = rows.Scan(&p.Slug, &p.Title, &p.Content, &p.CreatedAt, &p.Popularity)
		if err != nil {
			return []*previewing.Publication{}, err
		}
		publications = append(publications, &p)
	}
	return
}

func (p *Provider) PopularPublicationsProvider() (publications []*previewing.Publication, err error) {
	rows, err := p.db.QueryContext(p.ctx, popularPublicationsQuery, 5)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p previewing.Publication
		err = rows.Scan(&p.Slug, &p.Title, &p.Content, &p.CreatedAt, &p.Popularity)
		if err != nil {
			return []*previewing.Publication{}, err
		}
		publications = append(publications, &p)
	}
	return
}
