package rendrer

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tpml string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tpml)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("erro parsing : ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
