package handler

import (
	"net/http"
	"section_4/pkg/rendrer"
)

func Home(w http.ResponseWriter, r *http.Request) {
	rendrer.RenderTemplate(w, "home-page.html")
}
