package handlers

import (
	"net/http"

	"github.com/ihtgoot/i_learn/Section_3/pkg/rendrer"
)

// handeler of home page
func Home(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, "home-page.html")
}

// handels about page
func About(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, "about-page.html")
}
