package router

import (
	"net/http"

	"github.com/egoholic/router/params"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type Router struct {
	root *Node
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, p *params.Params)

func New() *Router {
	return &Router{NewNode("")}
}

func (r *Router) Root() *Node {
	return r.root
}

func (r *Router) Handler(p *params.Params) *Handler {
	return r.Root().Handler(p, p.NewIterator())
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	p := params.New(req.URL.String(), req.Method, map[string][]string{})
	handler := r.Handler(p)
	if handler == nil {
		return
	}

	fn := handler.HandlerFunc()
	if fn != nil {
		(*fn)(w, req, p)
	}
}

type Handler struct {
	handlerFunc *HandlerFunc
	desription  string
}

func newHandler(fn HandlerFunc, desc string) *Handler {
	return &Handler{&fn, desc}
}

func (h *Handler) HandlerFunc() *HandlerFunc {
	return h.handlerFunc
}

func (h *Handler) Description() string {
	return h.desription
}

type Node struct {
	pathChunk    string
	children     map[string]*Node
	verbHandlers map[string]*Handler
}

func NewNode(chunk string) *Node {
	return &Node{chunk, map[string]*Node{}, map[string]*Handler{}}
}

func (n *Node) Sub(chunk string) *Node {
	var node *Node
	node = n.children[chunk]
	if node != nil {
		return node
	}

	node = NewNode(chunk)
	n.children[chunk] = node
	return node
}

func (n *Node) Handler(p *params.Params, iter *params.PathChunksIterator) *Handler {
	if iter.HasNext() {
		chunk, _ := iter.Next()
		if child, ok := n.children[chunk]; ok {
			return child.Handler(p, iter)
		}
		return nil
	}
	return n.verbHandlers[p.Verb()]
}

func (n *Node) GET(fn HandlerFunc, d string) {
	n.verbHandlers[GET] = newHandler(fn, d)
}

func (n *Node) POST(fn HandlerFunc, d string) {
	n.verbHandlers[POST] = newHandler(fn, d)
}

func (n *Node) PUT(fn HandlerFunc, d string) {
	n.verbHandlers[PUT] = newHandler(fn, d)
}

func (n *Node) PATCH(fn HandlerFunc, d string) {
	n.verbHandlers[PATCH] = newHandler(fn, d)
}

func (n *Node) DELETE(fn HandlerFunc, d string) {
	n.verbHandlers[DELETE] = newHandler(fn, d)
}
