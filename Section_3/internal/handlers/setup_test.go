package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/ihtgoot/i_learn/Section_3/internal/config"
	"github.com/ihtgoot/i_learn/Section_3/internal/models"
	"github.com/ihtgoot/i_learn/Section_3/internal/rendrer"
	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	fmt.Println(session)
	return session.LoadAndSave(next)
}

// return template cache : created a map and stores the ample in for caching.
func TempleteCache() (map[string]*template.Template, error) {

	var Cacahe = make(map[string]*template.Template)

	// get all files *-page.html from folder ./template
	pages, err := filepath.Glob(fmt.Sprintf("%s*-page.tpml", parseTemplate))
	if err != nil {
		fmt.Println("errro in cacahing", err)
	}

	// range through kine of *-page.tpml
	for _, page := range pages {
		// page
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return Cacahe, err
		}
		//base page
		matche, err := filepath.Glob(fmt.Sprintf("%s*-layout.tpml", parseTemplate))
		if err != nil {
			return Cacahe, err
		}
		// use base
		if len(matche) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s*-layout.tpml", parseTemplate))
			if err != nil {
				return Cacahe, err
			}
		}
		// adding the page to cacahe
		Cacahe[name] = templateSet
	}
	return Cacahe, nil
}

var app config.AppConfig
var session *scs.SessionManager
var parseTemplate = "../../templates/"

func getRoutes() http.Handler {

	// data to be available in the session
	gob.Register(models.Reservation{})
	// change it to production when in use

	app.InProduction = true

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 2 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// create template static cache
	templeteCache, err := TempleteCache()
	if err != nil {
		log.Fatalln("error creating template cacahe ", err)
	} else {
		fmt.Println("cache made")
	}

	// string cacahe in cacahe but using config
	app.TemplateCacahe = templeteCache
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)
	rendrer.NewTemplate(&app)
	fmt.Println("server is running")

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// mux.Get("/", labadaba)
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/eremite", Repo.Eremite)
	mux.Get("/couple", Repo.Couple)
	mux.Get("/family", Repo.Family)
	mux.Get("/reservation", Repo.Reservation)
	mux.Post("/reservation", Repo.POST_Reservation)
	mux.Post("/reservation-json", Repo.ReservationJSON)
	mux.Get("/make-reservation", Repo.Make_Reservation)
	mux.Post("/make-reservation", Repo.Make_ReservationPOST)
	mux.Get("/reservation-overview", Repo.ReservationOverview)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
