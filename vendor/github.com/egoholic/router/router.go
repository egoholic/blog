package router

import (
	"net/http"

	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/node"
	"github.com/egoholic/router/params"
)

type Router struct {
	root *node.Node
}

func New() *Router {
	return &Router{node.New("", &node.DumbForm{})}
}
func (r *Router) Root() *node.Node {
	return r.root
}
func (r *Router) Handler(p *params.Params) *handler.Handler {
	return r.Root().Handler(p, p.NewIterator())
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	p := params.New(req.URL.String(), req.Method, map[string]interface{}{})
	handler := r.Handler(p)
	if handler == nil {
		return
	}
	fn := handler.HandlerFn()
	if fn != nil {
		fn(w, req, p)
	}
}
