package http

import (
	server "net/http"

	"github.com/egoholic/blog/blog/previewing"
)

type HTTPTransport struct {
	w server.ResponseWriter
}

func (t *HTTPTransport) Deliver(value *previewing.Value) (err error) {

}
