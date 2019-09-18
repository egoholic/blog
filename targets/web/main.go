package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	blogPreviewing "github.com/egoholic/blog/blog/previewing/handler/http"
	. "github.com/egoholic/blog/config"
	rtr "github.com/egoholic/router"
	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/params"
)

type HandlerFnBuilder func(context.Context, *sql.DB, *log.Logger) func(w http.ResponseWriter, r *http.Request, p *params.Params)

var (
	logger  = log.New(os.Stdout, "blog", 0)
	connStr string
	db      *sql.DB
	err     error
)

func main() {
	logger.Println("server starting...")
	connStr, err = Config.DBCredentials().ConnectionString()
	if err != nil {
		logger.Fatalf("ERROR1: %s\n", err.Error())
	}
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatalf("ERROR2: %s\n", err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		logger.Fatalf("ERROR3: %s\n", err.Error())
	}

	router := rtr.New()
	root := router.Root()
	root.GET(prepare(blogPreviewing.New), "presents recent and popular publications")

	logger.Println("server listens :3000 port")
	logger.Fatal(http.ListenAndServe(":3000", router))
}

func prepare(hb HandlerFnBuilder) handler.HandlerFn {
	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		d := 100 * time.Millisecond
		ctx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()
		h := hb(ctx, db, logger)
		h(w, r, p)
	}
}
