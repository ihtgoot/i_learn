package handlers

import (
	"main/pkg/rendrer"
	"net/http"
)

// handeler of home page
func Home(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, "home-page.html")
}

// handels about page
func About(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, "about-page.html")
}
