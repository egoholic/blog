package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	authorPreviewing "github.com/egoholic/blog/author/previewing/handler/http"
	blogPreviewing "github.com/egoholic/blog/blog/previewing/handler/http"
	"github.com/egoholic/blog/config"
	publicationReading "github.com/egoholic/blog/publication/reading/handler/http"
	rubricPreviewing "github.com/egoholic/blog/rubric/previewing/handler/http"

	. "github.com/egoholic/blog/config"
	rtr "github.com/egoholic/router"
	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/node"
	"github.com/egoholic/router/params"
)

type HandlerFnBuilder func(context.Context, *sql.DB, *log.Logger, func(http.ResponseWriter, *http.Request, *params.Params)) func(http.ResponseWriter, *http.Request, *params.Params)

var (
	logger       = log.New(LogFile, "blog", 0)
	connStr      string
	db           *sql.DB
	err          error
	notFoundView = template.Must(template.ParseFiles("shared/layouts/layout.html", "shared/templates/not_found.html"))
)

type SingleStringParamURLForm struct{}

func (f *SingleStringParamURLForm) CheckAndPopulate(pattern string, chunk string, prms *params.Params) bool {
	prms.Set(pattern, chunk)
	return true
}

func main() {
	// TODO: separate configs for different executables
	if config.Domain == "" {
		panic("No Blog Domain provided!")
	}
	logger.Println("server starting...")
	db, err = sql.Open("postgres", DBConnectionString)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	router := rtr.New()
	root := router.Root()
	// blog
	root.GET(prepare(blogPreviewing.New), "presents recent and popular publications")
	// publications
	publications := root.Child("p", &node.DumbForm{})
	publications.Child(":slug", &SingleStringParamURLForm{}).GET(prepare(publicationReading.New), "presents selected publications")
	// authors
	authors := root.Child("a", &node.DumbForm{})
	authors.Child(":login", &SingleStringParamURLForm{}).GET(prepare(authorPreviewing.New), "presents authors's bio and previews for  his/her publications")
	// rubrics
	rubrics := root.Child("r", &node.DumbForm{})
	rubrics.Child(":slug", &SingleStringParamURLForm{}).GET(prepare(rubricPreviewing.New), "presents rubrics and related publications")

	pid := os.Getpid()
	pidf, err := os.Create(PIDFilePath)
	if err != nil {
		fmt.Printf("FATALPID: %s   = %d\n", err.Error(), pid)
	}
	_, err = pidf.WriteString(strconv.Itoa(pid))
	if err != nil {
		fmt.Printf("FATALPID: %s   = %d\n", err.Error(), pid)
	}
	defer func() {
		err = pidf.Close()
		if err != nil {
			fmt.Printf("FATALPID: %s   = %d\n", err.Error(), pid)
		}
		err = os.Remove(PIDFilePath)
		if err != nil {
			fmt.Printf("FATALPID: %s   = %d\n", err.Error(), pid)
		}
	}()

	logger.Printf("server listens :%d port\n", Port)
	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), router))
}

func prepare(hb HandlerFnBuilder) handler.HandlerFn {
	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		d := 100 * time.Millisecond
		ctx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()
		h := hb(ctx, db, logger, notFound)
		h(w, r, p)
	}
}

func notFound(w http.ResponseWriter, r *http.Request, p *params.Params) {
	err := notFoundView.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
