package main

import (
	"net/http"

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

}

func showHome(w http.ResponseWriter, r *http.Request, p *params.Params) {

}
