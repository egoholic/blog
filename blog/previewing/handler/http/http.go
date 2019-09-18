package http

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/egoholic/blog/blog/previewing"
	repository "github.com/egoholic/blog/blog/previewing/repository/postgresql"
	"github.com/egoholic/router/params"
)

var view = template.Must(template.ParseFiles("blog/previewing/handler/http/templates/main.html", "shared/layout/layout.html"))

func New(ctx context.Context, db *sql.DB, logger *log.Logger) func(w http.ResponseWriter, r *http.Request, p *params.Params) {
	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		logger.Println("handles blog preview")
		repo := repository.New(ctx, db, logger)
		value := previewing.New(logger, repo, repo)
		//content := layout.New(value, "PageTitle", "BlogTitle", "Blog description.", "blog;keywords")
		value.PopularPublications()
		value.RecentPublications()
		err := view.ExecuteTemplate(w, "layout.html", value)
		//err := view.Execute(w, value)
		if err != nil {
			logger.Panicf("ERROR-http.go: %s\n", err.Error())
		}

		//w.WriteHeader(200)
	}
}
