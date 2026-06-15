package main

import "net/http"

func main() {
	s := &api{
		addr: ":8000",
	}
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", s.getUserHandler)
	mux.HandleFunc("POST /users", s.createUserHandler)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
