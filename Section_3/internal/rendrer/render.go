package rendrer

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ihtgoot/i_learn/Section_3/internal/config"
	"github.com/ihtgoot/i_learn/Section_3/internal/models"
	"github.com/justinas/nosurf"
)

// // render templete server as a wrapper and a reader
// // a layout and a template from folder /templete to describe writer

// func RenderTemplateTempV0(w http.ResponseWriter, tpml string) {
// 	parseTemplate, err := template.ParseFiles("/home/ihtgoot/web_dev/Section_3/templates/"+tpml, "/home/ihtgoot/web_dev/Section_3/templates/base-layout.tpml")
// 	if err != nil {
// 		fmt.Println("Error parsing templates:", err)
// 		http.Error(w, "Internal Server Error 1 ", http.StatusInternalServerError)
// 		return
// 	}
// 	err = parseTemplate.ExecuteTemplate(w, "base", nil)
// 	if err != nil {
// 		fmt.Println("error executing template : ", err)
// 		http.Error(w, "Internal Server Error 2 ", http.StatusInternalServerError)
// 		return
// 	}
// }

// // custm template cache
// var templeteCacheV0 = make(map[string]*template.Template) //this is our cache : any time we can seach out template and it return the whole content
// func RenderTemplateCacheDynamic(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check we there if we have already cached out templete : whast in cache
// 	_, inMap := templeteCacheV0[t] // checks weather something like t in in cache
// 	if !inMap {
// 		// need to create templete
// 		err = createTempleteCacheV0(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// already templete in out cache we use it
// 		log.Println("using from cache")
// 	}

// 	tmpl = templeteCacheV0[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// // this is the simplest solution ; wors well if small no of files
// // it grows as the no of pages increase and it never is cleared and since we hardcode path we need to add new temp manually
// // error would only be discovered on second run (on cache access)
// // this is efficeint not good
// func createTempleteCacheV0(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("/home/ihtgoot/web_dev/Section_3/templates/%s", t),
// 		"/home/ihtgoot/web_dev/Section_3/templates/base-layout.tpml",
// 	}
// 	// parse template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// add templete to cache map
// 	templeteCacheV0[t] = tmpl

// 	return nil
// }

// // satatic cache : on every time application caceh is build and it holds everything that ends with .html or .tpml
// // easieer to mainatian and felxible to adjustment
// // caceh is available everytime on startup entirely , nothing has to be build everythime thus reducing responase time
// // Disadvantage : change to conetent is not immediately visible : we can have a turaround by making a config file to disable cache

// // AddDefault contains DATA which will be added to data sent to template
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

var app *config.AppConfig
var parseTemplates = "./templates/"

func NewRendrer(a *config.AppConfig) {
	app = a
}

func Template(w http.ResponseWriter, r *http.Request, tpml string, td *models.TemplateData) error {

	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCacahe
	} else {
		templateCache, _ = CreateTempleteCache()
	}

	// create template static cache
	// templeteCache, err := CreateTempleteCache()
	// if err != nil {
	// 	log.Fatalln("error creating template cacahe ", err)
	// }

	t, ok := templateCache[tpml]
	if !ok {
		return errors.New("template not in cacahe for some reason")
	}

	// check if t has a valid template ;
	// store the resualt of t in a buffer and double-check if it is valid value
	buff := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := t.Execute(buff, td)
	if err != nil {
		log.Println(err)
		return err
	}

	//render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		fmt.Println("error parsing ", err)
		return err
	}

	//parseTemplate, err := template.ParseFiles("/home/ihtgoot/web_dev/Section_3/templates/"+tpml, "/home/ihtgoot/web_dev/Section_3/templates/base-layout.html")
	//if err != nil {
	//	fmt.Println("errro parsing template ", err)
	//}
	//err = parseTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("errro executing template", err)
	//}
	if err != nil {
		return err
	}
	return nil
}

// return template cache : created a map and stores the ample in for caching.
func CreateTempleteCache() (map[string]*template.Template, error) {

	var Cacahe = make(map[string]*template.Template)

	// get all files *-page.html from folder ./template
	pages, err := filepath.Glob(fmt.Sprintf("%s*-page.tpml", parseTemplates))
	if err != nil {
		fmt.Println("errro in cacahing", err)
	}

	// range through kine of *-page.tpml
	for _, page := range pages {
		// page
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return Cacahe, err
		}
		//base page
		matche, err := filepath.Glob(fmt.Sprintf("%s*-layout.tpml", parseTemplates))
		if err != nil {
			return Cacahe, err
		}
		// use base
		if len(matche) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s*-layout.tpml", parseTemplates))
			if err != nil {
				return Cacahe, err
			}
		}
		// adding the page to cacahe
		Cacahe[name] = templateSet
	}
	return Cacahe, nil
}
