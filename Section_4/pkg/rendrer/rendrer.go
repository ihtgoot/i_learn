package rendrer

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tpml string) {
	pars, err := template.ParseFiles("/home/ihtgoot/web_dev/Section_4/templates/"+tpml, "/home/ihtgoot/web_dev/Section_4/templates/base-layout.html")
	if err != nil {
		fmt.Println("erro in parsing : ", err)
		http.Error(w, "error parsing : ", http.StatusInternalServerError)
		return
	}
	err = pars.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println("erro in executing : ", err)
		http.Error(w, "error parsing : ", http.StatusInternalServerError)
		return
	}
}
