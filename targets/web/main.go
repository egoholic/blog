package main

import (
	"net/http"

	"github.com/egoholic/blog/view"
	"github.com/egoholic/blog/view/homepage"
	rtr "github.com/egoholic/router"
	"github.com/egoholic/router/params"
)

var router *rtr.Router

func init() {
	router = rtr.New()
	root := router.Root()
	root.GET(showHome, "presents recent stuff")
}
func main() {
	http.ListenAndServe(":8080", router)
}

func showHome(w http.ResponseWriter, r *http.Request, p *params.Params) {
	meta := view.Meta{"Homepage", "blog's homepage", "blog, homepage"}
	header := view.Header{"My Blog", "Best blog in the Universe!"}
	footer := view.Footer{"-|-", "2019-2020 Blog (c)"}
	content := view.Content{"MAIN", "Rubric1, Rubric2"}
	hp := homepage.New(meta, header, content, footer)
	hp.Execute(w)
}
