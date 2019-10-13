package http

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/egoholic/blog/author/previewing"
	repository "github.com/egoholic/blog/author/previewing/repository/postgresql"
	"github.com/egoholic/router/params"
)

var view = template.Must(template.ParseFiles("shared/layouts/layout.html", "shared/views/publication/li.html", "author/previewing/handler/http/templates/content.html"))

func New(ctx context.Context, db *sql.DB, logger *log.Logger, notFound func(http.ResponseWriter, *http.Request, *params.Params)) func(w http.ResponseWriter, r *http.Request, p *params.Params) {
	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		repo := repository.New(ctx, db, logger)
		param, ok := p.Get(":login")
		if !ok {
			logger.Panicln("ERROR: empty login param")
		}
		login, ok := param.(string)
		if !ok {
			logger.Panicf("ERROR: login param should be string. (%#v) given\n", login)
		}
		value, err := previewing.New(logger, repo, repo, login)
		if err != nil {
			logger.Printf("ERROR: %s", err.Error())
			notFound(w, r, p)
			return
		}
		err = view.Execute(w, value)
		if err != nil {
			logger.Panicf("ERROR: %s\n", err.Error())
		}
		w.Header().Set("Content-Type", "text/html")
	}
}
