package webhooks

import (
	"net/http"
)

func NewListener(address string, handler Handler) *Listener {
	return &Listener{
		address: address,
		handler: handler,
	}
}

type Listener struct {
	address string
	handler Handler
}

func (i *Listener) ListenAndWait() error {
	serverErrChan := make(chan error)
	doneChan := make(chan bool)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		i.handler.HandleRequest(w, r)
		if i.handler.IsComplete() {
			doneChan <- true
		}
	})

	server := &http.Server{
		Addr:    i.address,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			serverErrChan <- err
		}
	}()

	select {
	case err := <-serverErrChan:
		return err
	case <-doneChan:
		return i.handler.Error()
	}
}
