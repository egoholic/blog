package postgresql

import (
	"context"
	"database/sql"
	"log"

	"github.com/egoholic/blog/author/previewing"
	_ "github.com/lib/pq"
)

type Repository struct {
	db     *sql.DB
	ctx    context.Context
	logger *log.Logger
}

var (
	authorQuery = `SELECT a.first_name || ' ' || a.last_name AS full_name,
												a.bio                              AS bio,
												a.login                            AS login
												FROM accounts AS a
												WHERE a.login = $1
												LIMIT 1;`

	publicationsQuery = `SELECT slug,
															title,
															created_at,
															popularity
											FROM publications AS p
											INNER JOIN publication_authors AS pa
												      ON pa.author_login     = $1
														 AND pa.publication_slug = p.slug
											ORDER BY created_at DESC;`
)

func New(ctx context.Context, db *sql.DB, logger *log.Logger) *Repository {
	return &Repository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

func (r *Repository) AuthorByLogin(l string) (*previewing.Author, error) {
	var a previewing.Author
	row := r.db.QueryRowContext(r.ctx, authorQuery, l)
	err := row.Scan(&a.FullName, &a.Bio, &a.Login)
	return &a, err
}

func (r *Repository) PublicationsOf(l string) (publications []*previewing.Publication, err error) {
	rows, err := r.db.QueryContext(r.ctx, publicationsQuery, l)
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
