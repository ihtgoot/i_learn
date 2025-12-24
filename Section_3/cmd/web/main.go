package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ihtgoot/i_learn/Section_3/internal/config"
	"github.com/ihtgoot/i_learn/Section_3/internal/handlers"
	"github.com/ihtgoot/i_learn/Section_3/internal/models"
	"github.com/ihtgoot/i_learn/Section_3/internal/rendrer"
)

var app config.AppConfig
var session *scs.SessionManager

const portNumber string = ":8080"

func main() {
	fmt.Println("start server")
	err := run()

	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	fmt.Println("we started on port number : ", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
	// http.ListenAndServe(portNumber, nil)
	os.Exit(0)
}

func run() error {
	// data to be available in the session
	gob.Register(models.Reservation{})

	// change it to production when in use
	app.InProduction = true

	session = scs.New()
	session.Lifetime = 1 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	// create template static cache
	templeteCache, err := rendrer.CreateTempleteCache()
	if err != nil {
		log.Fatalln("error creating template cacahe ", err)
		return err
	} else {
		fmt.Println("cache made")
	}

	// string cacahe in cacahe but using config
	app.TemplateCacahe = templeteCache
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	rendrer.NewTemplate(&app)

	fmt.Println("server is running")

	return nil
}

// Fprintf , write the string in a io buffer ,here the string is "hello"
// and the io buffer is the http.ResponseWriter i.e w

// on successful execution n stores an int value , stores the no. of bytes
// written to response writer, and err is left empty awhile if there is an
// error err store errro message

// in HandelFunc first is the dir of out server that is handles and
// second arguiment is the fnction that defines how it respaonse if the
//  dir is called ;
// techinically HandelFunc assign a handler to the dir on access firstis
// the dir and second is the function that handles the dir

// since go can have more than one return value it provides us with the
// functionality to handle these during operation without crashig in
// production this make for a simplifies apprch to error handeling

//  http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
//	n, err := fmt.Fprint(w, "hi")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf(fmt.Sprintf("\nBytes Writen :n : %d", n))
//})
