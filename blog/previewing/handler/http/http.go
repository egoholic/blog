package http

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/egoholic/blog/blog/previewing"
	repository "github.com/egoholic/blog/blog/previewing/repository/postgresql"
	"github.com/egoholic/router/params"
)

var view = template.Must(template.ParseFiles("shared/layout/layout.html", "blog/previewing/handler/http/templates/content.html"))

func New(ctx context.Context, db *sql.DB, logger *log.Logger) func(w http.ResponseWriter, r *http.Request, p *params.Params) {
	logger.Println("handles blog preview")

	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		repo := repository.New(ctx, db, logger)
		value := previewing.New(logger, repo, repo, repo)
		fmt.Printf("\n\n\n%#v\n\n\n", value)
		pp, err := repo.PopularPublications()
		logger.Println("popular-publications", pp, err)
		fmt.Println("popular-publications", pp, err)
		rp, err := repo.RecentPublications()
		logger.Println("recent-publications", rp, err)
		fmt.Println("recent-publications", rp, err)

		rb, err := repo.Rubrics()
		logger.Println("rubrics", rb, err)
		fmt.Println("rubrics", rb, err)

		err = view.Execute(w, value)
		if err != nil {
			logger.Panicf("ERROR-http.go: %s\n", err.Error())
		}
		w.Header().Set("Content-Type", "text/html")
	}
}
