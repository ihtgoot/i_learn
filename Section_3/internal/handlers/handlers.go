package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ihtgoot/i_learn/Section_3/internal/config"
	"github.com/ihtgoot/i_learn/Section_3/internal/models"
	"github.com/ihtgoot/i_learn/Section_3/internal/rendrer"
)

// Reepository is the repository type
type Repository struct {
	App *config.AppConfig
}

// repo the repository used by the handler
var Repo *Repository

// newwwRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New Handlers set the repsitory for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// handeler of home page
func (h *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// store remote IP of user
	remoteIP := r.RemoteAddr
	h.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	rendrer.RenderTemplate(w, r, "home-page.tpml", &models.TemplateData{})
}

// handels about page
func (a *Repository) About(w http.ResponseWriter, r *http.Request) {
	//a.App.Session
	// sidekickmap := make(map[string]string)
	// sidekickmap["morty"] = "ooh , wee"
	// remoteIP := a.App.Session.GetString(r.Context(), "remote_ip")
	// sidekickmap["remote_ip"] = remoteIP
	rendrer.RenderTemplate(w, r, "about-page.tpml", &models.TemplateData{})
}

// handels Contact page
func (a *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, r, "contact-page.tpml", &models.TemplateData{})
}

// handels Erermite page
func (a *Repository) Eremite(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, r, "eremite-page.tpml", &models.TemplateData{})
}

// handels Couple page
func (a *Repository) Couple(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, r, "couple-page.tpml", &models.TemplateData{})
}

// handels Family page
func (a *Repository) Family(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, r, "family-page.tpml", &models.TemplateData{})
}

// handels Reservation page GET request
func (a *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, r, "check-availability-page.tpml", &models.TemplateData{})
}

// handels Reservation page POST request
func (a *Repository) POST_Reservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("startingDate")
	end := r.Form.Get("endingDate")
	fmt.Println(start, end)
	w.Write([]byte(fmt.Sprintf("Success you ucking send something, start : %s, end : %s", start, end)))
}

// handels resrervation json it retur s json (json : javascript obect resource)

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"messgae`
}

func (a *Repository) ReservationJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Its available",
	}
	out, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Println(err)
	}
	// log.Println(string(out))
	w.Header().Set("content-Type", "application/json") // tell browser what the dtatype is o that it has a header
	w.Write(out)
}

// handels make-reservation page
func (a *Repository) Make_Reservation(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, r, "make-reservation-page.tpml", &models.TemplateData{})
}

// repo pattern : make use of packages across applications , this hellps us in developement : we can centrally turn our cache on/off

// ❯ go run ./
// package github.com/ihtgoot/i_learn/Section_3/cmd/web
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/handlers
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/rendrer
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/handlers: import cycle not allowed
