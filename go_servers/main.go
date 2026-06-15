package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) handleUserMethods(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome user"))
}
func (s *api) handleAdminMethods(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome admin"))
}

func main() {
	s := &api{
		addr: ":8000",
	}
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", s.handleUserMethods)
	mux.HandleFunc("GET /admins", s.handleAdminMethods)
	srv.ListenAndServe()

}
