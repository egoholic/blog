package http

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/egoholic/blog/rubric/previewing"
	repository "github.com/egoholic/blog/rubric/previewing/repository/postgresql"
	"github.com/egoholic/router/params"
)

var view = template.Must(template.ParseFiles("shared/layout/layout.html", "rubric/previewing/handler/http/templates/content.html"))

func New(ctx context.Context, db *sql.DB, logger *log.Logger) func(w http.ResponseWriter, r *http.Request, p *params.Params) {
	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		repo := repository.New(ctx, db, logger)
		param, ok := p.Get(":slug")
		if !ok {
			logger.Panicln("ERROR: empty slug param")
		}
		slug, ok := param.(string)
		if !ok {
			logger.Panicf("ERROR: slug param should be string. (%#v) given\n", slug)
		}
		value := previewing.New(logger, repo, repo, slug)
		err := view.Execute(w, value)
		if err != nil {
			logger.Panicf("ERROR: %s\n", err.Error())
		}
		w.Header().Set("Content-Type", "text/html")
	}
}
