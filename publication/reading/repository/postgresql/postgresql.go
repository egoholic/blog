package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/egoholic/blog/publication/reading"
	_ "github.com/lib/pq"
)

type Repository struct {
	db     *sql.DB
	ctx    context.Context
	logger *log.Logger
}

var (
	blogQuery = `SELECT title
							 FROM (SELECT domain,
														title
										 FROM blogs
										 WHERE domain = $1
										 LIMIT 1) AS b;`
	publicationQuery = `SELECT  slug,
															title,
															content,
															created_at,
															popularity,
															meta_keywords,
															meta_description
											FROM publications
											WHERE slug = $1;`

	authorsQuery = `SELECT a.first_name || ' ' || a.last_name AS full_name,
												 a.bio                              AS bio,
												 a.login                            AS login
												 FROM       accounts            AS a
												 INNER JOIN publication_authors AS pa
												         ON pa.author_login     = a.login
												        AND pa.publication_slug = $1;`
)

func New(ctx context.Context, db *sql.DB, logger *log.Logger) *Repository {
	return &Repository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

func (r *Repository) PublicationBySlug(s string) (*reading.Publication, error) {
	var p reading.Publication
	row := r.db.QueryRowContext(r.ctx, publicationQuery, s)
	err := row.Scan(&p.Slug, &p.Title, &p.Content, &p.CreatedAt, &p.Popularity, &p.MetaKeywords, &p.MetaDescription)
	return &p, err
}

func (r *Repository) AuthorsOf(s string) (authors []*reading.Author, err error) {
	rows, err := r.db.QueryContext(r.ctx, authorsQuery, s)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var a reading.Author
		err = rows.Scan(&a.FullName, &a.Bio, &a.Login)
		if err != nil {
			return
		}
		authors = append(authors, &a)
	}
	return
}
func (r *Repository) BlogByDomain(domain string) (*reading.Blog, error) {
	var blog reading.Blog
	row := r.db.QueryRowContext(r.ctx, blogQuery, domain)
	err := row.Scan(&blog.Title)
	if err != nil {
		fmt.Printf("publication-reading repo err: %s", err.Error())
	}
	return &blog, err
}
