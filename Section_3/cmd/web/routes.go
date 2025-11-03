package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/ihtgoot/i_learn/Section_3/pkg/config"
	"github.com/ihtgoot/i_learn/Section_3/pkg/handlers"
	"github.com/ihtgoot/i_learn/Section_3/pkg/helper"
)

func labadaba(w http.ResponseWriter, r *http.Request) {
	var owner, saying string
	owner, saying = helper.Getdata()
	n, err := fmt.Fprintf(w, fmt.Sprintf("<html>this is the about page <br> %s said \" %s \" times</html>", owner, saying, addValues(4949, 5948)))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(fmt.Sprintf("\nBytes Written: n : %d", n))
}

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurd)
	mux.Use(SessionLoad)

	mux.Get("/", labadaba)
	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
