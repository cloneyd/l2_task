package chainOfResponisibility

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type HTTPMiddleware func(w http.ResponseWriter, r *http.Request, next http.Handler)

func AuthorizationMiddleware(w http.ResponseWriter, r *http.Request, next http.Handler) {
	//Authorizing
	next.ServeHTTP(w, r)
}

func PrivatePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You successfully logged in and accessed this private page!")
}

func MiddlewareToHandler(middleware HTTPMiddleware, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware(w, r, next)
	})
}

func StartServer() {
	ts := httptest.NewServer(MiddlewareToHandler(AuthorizationMiddleware, PrivatePageHandler))
	defer ts.Close()
}
