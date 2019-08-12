package node

import (
	"github.com/egoholic/router/handler"
	"github.com/egoholic/router/params"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type Form interface {
	CheckAndPopulate(pattern string, pathChunk string, prms *params.Params) bool
}
type Node struct {
	pathChunk       string
	form            Form
	staticChildren  map[string]*Node
	dynamicChildren map[string]*Node
	verbHandlers    map[string]*handler.Handler
}

func New(pch string, form Form) *Node {
	return &Node{
		pathChunk:       pch,
		form:            form,
		staticChildren:  map[string]*Node{},
		dynamicChildren: map[string]*Node{},
		verbHandlers:    map[string]*handler.Handler{},
	}
}
func (n *Node) Child(pathChunk string, form Form) *Node {
	var (
		node *Node
		ok   bool
	)
	if pathChunk[0] == ':' {
		node, ok = n.dynamicChildren[pathChunk]
		if ok {
			return node
		}
		node = New(pathChunk, form)
		n.dynamicChildren[pathChunk] = node
		return node
	}
	node, ok = n.staticChildren[pathChunk]
	if ok {
		return node
	}
	node = New(pathChunk, form)
	n.staticChildren[pathChunk] = node
	return node
}
func (n *Node) Handler(prms *params.Params, pathChunks *params.PathChunksIterator) *handler.Handler {
	if n.form.CheckAndPopulate(n.pathChunk, pathChunks.Current(), prms) {
		if pathChunks.HasNext() {
			chunk, _ := pathChunks.Next()
			child, ok := n.staticChildren[chunk]
			if ok && child.form.CheckAndPopulate(child.pathChunk, pathChunks.Current(), prms) {
				return child.Handler(prms, pathChunks)
			}
			for _, child := range n.dynamicChildren {
				if child != nil && child.form.CheckAndPopulate(child.pathChunk, pathChunks.Current(), prms) {
					return child.Handler(prms, pathChunks)
				}
			}
		} else {
			return n.verbHandlers[prms.Verb()]
		}
	}
	return nil
}
func (n *Node) GET(fn handler.HandlerFn, d string) {
	n.verbHandlers[GET] = handler.New(fn, d)
}
func (n *Node) POST(fn handler.HandlerFn, d string) {
	n.verbHandlers[POST] = handler.New(fn, d)
}
func (n *Node) PUT(fn handler.HandlerFn, d string) {
	n.verbHandlers[PUT] = handler.New(fn, d)
}
func (n *Node) PATCH(fn handler.HandlerFn, d string) {
	n.verbHandlers[PATCH] = handler.New(fn, d)
}
func (n *Node) DELETE(fn handler.HandlerFn, d string) {
	n.verbHandlers[DELETE] = handler.New(fn, d)
}

type DumbForm struct{}

func (_ *DumbForm) CheckAndPopulate(_ string, _ string, _ *params.Params) bool {
	return true
}
