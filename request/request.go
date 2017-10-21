package request

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const IdHeaderName = "X-Request-Id"

type HttpRequest struct {
	*http.Request
	Params httprouter.Params
	Id     string
}
