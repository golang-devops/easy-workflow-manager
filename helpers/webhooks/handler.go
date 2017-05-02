package webhooks

import (
	"net/http"
)

type Handler interface {
	HandleRequest(writer http.ResponseWriter, req *http.Request)
	IsComplete() bool
	Error() error
}
