package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func hitLogger(next http.Handler) http.Handler {
	//useless
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("working")
		next.ServeHTTP(w, r)
	})
}

// NoSurd server as a CSRF protection middleware
func NoSurd(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	fmt.Println(csrfHandler)
	return csrfHandler
}

// saves session data for each request
func SessionLoad(next http.Handler) http.Handler {
	fmt.Println(session)
	return session.LoadAndSave(next)
}
