package main

import (
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		}
	default:
		w.Write([]byte("404 not found"))
		return
	}
}

func smain() {
	s := &server{
		addr: ":8000",
	}
	http.ListenAndServe(s.addr, s)
}
