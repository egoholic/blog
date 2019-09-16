package main

import (
	"log"
	"net/http"
	"os"

	rtr "github.com/egoholic/router"
	"github.com/egoholic/router/params"
)

var router *rtr.Router
var l = log.New(os.Stdout, "blog", 0)

func init() {
	router = rtr.New()
	root := router.Root()
	root.GET(showHome, "presents recent stuff")
}
func main() {
	l.Println("server listens to :3000 port")
	l.Fatal(http.ListenAndServe(":3000", router))
}

func showHome(w http.ResponseWriter, r *http.Request, p *params.Params) {
	// d := hp.Destination(w)
	// prv := &previewing.Value{}
	// prv.Deliver(nil, nil, d)
	l.Println(r.Header)
}
