package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/ihtgoot/i_learn/Section_3/internal/config"
	"github.com/ihtgoot/i_learn/Section_3/internal/handlers"
	"github.com/ihtgoot/i_learn/Section_3/internal/helper"
	"github.com/justinas/nosurf"
)

func labadaba(w http.ResponseWriter, r *http.Request) {
	var owner, saying string

	token := nosurf.Token(r)

	// fmt.Fprintf(w, `
	// 	<form method="POST" action="/submit">
	// 		<input type="hidden" name="csrf_token" value="%s">
	// 		<button type="submit">Send</button>
	// 	</form>
	// `, token)
	fmt.Println(token)

	owner, saying = helper.Getdata()
	html := fmt.Sprintf("<html>this is the about page <br> %s said \" %s \" times</html>", owner, saying)
	n, err := fmt.Fprint(w, html)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nBytes Written: n : %d", n)
}

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(hitLogger)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// mux.Get("/", labadaba)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/eremite", handlers.Repo.Eremite)
	mux.Get("/couple", handlers.Repo.Couple)
	mux.Get("/family", handlers.Repo.Family)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Post("/reservation", handlers.Repo.POST_Reservation)
	mux.Post("/reservation-json", handlers.Repo.ReservationJSON)
	mux.Get("/choose-banglow/{id}", handlers.Repo.ChooseBanglow)
	mux.Get("/book-banglow", handlers.Repo.BookBanglow)
	mux.Get("/make-reservation", handlers.Repo.Make_Reservation)
	mux.Post("/make-reservation", handlers.Repo.Make_ReservationPOST)
	mux.Get("/reservation-overview", handlers.Repo.ReservationOverview)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
