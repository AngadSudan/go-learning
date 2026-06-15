package main

import (
	"fmt"
	"megaproject/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":8000"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	fmt.Println("Server running at port ", cfg.addr)
	app.run(mux)

}
