package http

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/egoholic/blog/blog/previewing"
	repository "github.com/egoholic/blog/blog/previewing/repository/postgresql"
	"github.com/egoholic/blog/config"
	"github.com/egoholic/router/params"
)

var view = template.Must(template.ParseFiles("shared/layouts/layout.html", "shared/views/publication/li.html", "shared/views/rubric/li.html", "blog/previewing/handler/http/templates/content.html"))

func New(ctx context.Context, db *sql.DB, logger *log.Logger, not_found func(http.ResponseWriter, *http.Request, *params.Params)) func(w http.ResponseWriter, r *http.Request, p *params.Params) {
	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		repo := repository.New(ctx, db, logger)
		value, err := previewing.New(logger, config.Domain, repo, repo, repo, repo)
		if err != nil {
			logger.Printf("ERROR: %s", err.Error())
			not_found(w, r, p)
		}
		err = view.Execute(w, value)
		if err != nil {
			logger.Panic(err.Error())
		}
		w.Header().Set("Content-Type", "text/html")
	}
}
