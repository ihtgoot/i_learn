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
	"github.com/ihtgoot/i_learn/Section_3/internal/driver"
	"github.com/ihtgoot/i_learn/Section_3/internal/handlers"
	"github.com/ihtgoot/i_learn/Section_3/internal/helper"
	"github.com/ihtgoot/i_learn/Section_3/internal/models"
	"github.com/ihtgoot/i_learn/Section_3/internal/rendrer"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

const portNumber string = ":8080"

func main() {
	fmt.Println("start server")
	db, err := run()
	defer db.SQl.Close()
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

func run() (*driver.DB, error) {
	// data to be available in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Banglow{})
	gob.Register(models.Restriction{})
	gob.Register(models.BanglowRestriction{})
	// “Gob does not handle user input or database communication.
	// It is only used by the session system to encode and decode
	// Go structs so data can persist between pages and API endpoints.
	// By calling gob.Register, we tell gob how to interpret
	// the concrete structure of values stored in the session.”
	// Gob freezes Go values into bytes and later restores them exactly as they were.

	// change it to production when in use
	app.InProduction = true

	infoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 1 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connecting to database
	log.Println("connecting to db .................")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=web_dev user=postgres password= 1729@Iwilldoit") // libpq connection string
	if err != nil {
		log.Fatal("No connection to database Terminating .................... ")
	}
	log.Println("Successfully connectde to database")

	// create template static cache
	templeteCache, err := rendrer.CreateTempleteCache()
	if err != nil {
		log.Fatalln("error creating template cacahe ", err)
		return nil, err
	} else {
		fmt.Println("cache made")
	}

	// string cacahe in cacahe but using config
	app.TemplateCacahe = templeteCache
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helper.NewHelper(&app)
	rendrer.NewRendrer(&app)

	fmt.Println("server is running")

	return db, nil
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
