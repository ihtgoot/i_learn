package rendrer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// render templete server as a wrapper and a reader
// a layout and a template from folder /templete to describe writer

func RenderTemplateTemp(w http.ResponseWriter, tpml string) {
	parseTemplate, err := template.ParseFiles("/home/ihtgoot/web_dev/Section_3/templates/"+tpml, "/home/ihtgoot/web_dev/Section_3/templates/base-layout.html")
	if err != nil {
		fmt.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error 1 ", http.StatusInternalServerError)
		return
	}
	err = parseTemplate.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println("error executing template : ", err)
		http.Error(w, "Internal Server Error 2 ", http.StatusInternalServerError)
		return
	}
}

var templeteCache = make(map[string]*template.Template) //this is our cache : any time ewe can sefch out template and it return the whole content

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check wether if we have already cached out templete : whast in cache
	_, inMap := templeteCache[t] // checks weather something like t in in cache
	if !inMap {
		// need to create templete
		err = createTempleteCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// already templete in out cache we use it
		log.Println("using from cache")
	}

	tmpl = templeteCache[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// this is the simplest solution ; wors well if small no of files ;
//
//	it grows as the no of pages increase and it never is cleared and since we hardcode path we need to add new temp manually ;
//
// error would only be discovered on second run (on cache access) ;
//
//	this is efficeint not good
func createTempleteCache(t string) error {
	templates := []string{
		fmt.Sprintf("/home/ihtgoot/web_dev/Section_3/templates/%s", t),
		"/home/ihtgoot/web_dev/Section_3/templates/base-layout.html",
	}
	// parse template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add templete to cache map
	templeteCache[t] = tmpl

	return nil
}
