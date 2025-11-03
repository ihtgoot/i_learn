package handlers

import (
	"net/http"

	"github.com/ihtgoot/i_learn/Section_3/pkg/config"
	"github.com/ihtgoot/i_learn/Section_3/pkg/models"
	"github.com/ihtgoot/i_learn/Section_3/pkg/rendrer"
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

	rendrer.RenderTemplate(w, "home-page.html", &models.TemplateData{})
}

// handels about page
func (a *Repository) About(w http.ResponseWriter, r *http.Request) {
	//a.App.Session
	sidekickmap := make(map[string]string)
	sidekickmap["morty"] = "ooh , wee"
	remoteIP := a.App.Session.GetString(r.Context(), "remote_ip")
	sidekickmap["remote_ip"] = remoteIP
	rendrer.RenderTemplate(w, "about-page.html", &models.TemplateData{
		StringMap: sidekickmap,
	})
}

// repo pattern : make use of packages across applications , this hellps us in developement : we can centrally turn our cache on/off

// ❯ go run ./
// package github.com/ihtgoot/i_learn/Section_3/cmd/web
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/handlers
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/rendrer
//         imports github.com/ihtgoot/i_learn/Section_3/pkg/handlers: import cycle not allowed
