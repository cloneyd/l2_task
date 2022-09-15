package command

import (
	"fmt"
	"net/http"
)

// Command
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (hf HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hf(w, r)
}

// Concrete command
func RootHandler() Handler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you're requesting URI: %s\n", r.RequestURI)
	})
}

func StartServer() {
	// Invoker
	s := http.Server{
		Addr:    ":8080",
		Handler: RootHandler(),
	}
	s.ListenAndServe()
}
