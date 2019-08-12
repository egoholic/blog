package handler

import (
	"net/http"

	"github.com/egoholic/router/params"
)

type HandlerFn func(w http.ResponseWriter, r *http.Request, p *params.Params)

type Handler struct {
	handlerFn  HandlerFn
	desription string
}

func New(fn HandlerFn, desc string) *Handler {
	return &Handler{fn, desc}
}
func (h *Handler) HandlerFn() HandlerFn {
	return h.handlerFn
}
func (h *Handler) Description() string {
	return h.desription
}
