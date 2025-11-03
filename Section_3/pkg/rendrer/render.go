package rendrer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ihtgoot/i_learn/Section_3/pkg/config"
	"github.com/ihtgoot/i_learn/Section_3/pkg/models"
)

// render templete server as a wrapper and a reader
// a layout and a template from folder /templete to describe writer

func RenderTemplateTempV0(w http.ResponseWriter, tpml string) {
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

// custm template cache
var templeteCacheV0 = make(map[string]*template.Template) //this is our cache : any time we can seach out template and it return the whole content
func RenderTemplateCacheDynamic(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check we there if we have already cached out templete : whast in cache
	_, inMap := templeteCacheV0[t] // checks weather something like t in in cache
	if !inMap {
		// need to create templete
		err = createTempleteCacheV0(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// already templete in out cache we use it
		log.Println("using from cache")
	}

	tmpl = templeteCacheV0[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// this is the simplest solution ; wors well if small no of files
// it grows as the no of pages increase and it never is cleared and since we hardcode path we need to add new temp manually
// error would only be discovered on second run (on cache access)
// this is efficeint not good
func createTempleteCacheV0(t string) error {
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
	templeteCacheV0[t] = tmpl

	return nil
}

// satatic cache : on every time application caceh is build and it holds everything that ends with .html or .tpml
// easieer to mainatian and felxible to adjustment
// caceh is available everytime on startup entirely , nothing has to be build everythime thus reducing responase time
// Disadvantage : change to conetent is not immediately visible : we can have a turaround by making a config file to disable cache

// AddDefault contains DATA which will be added to data sent to template
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tpml string, td *models.TemplateData) {

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

	// get templateCache from appconfig

	// get the right template from cache
	t, ok := templateCache[tpml]
	if !ok {
		log.Fatalln("template not in cacahe for some reason\n", ok)
	}

	// check if t has a valid template ;
	// store the resualt of t in a buffer and double-check if it is valid value
	buff := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		fmt.Println("error parsing ", err)
	}

	//parseTemplate, err := template.ParseFiles("/home/ihtgoot/web_dev/Section_3/templates/"+tpml, "/home/ihtgoot/web_dev/Section_3/templates/base-layout.html")
	//if err != nil {
	//	fmt.Println("errro parsing template ", err)
	//}
	//err = parseTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("errro executing template", err)
	//}
}

// return template cache
func CreateTempleteCache() (map[string]*template.Template, error) {
	var Cacahe = make(map[string]*template.Template)
	// get all files *-page.html from folder ./template
	pages, err := filepath.Glob("/home/ihtgoot/web_dev/Section_3/templates/*-page.html")
	if err != nil {
		fmt.Println("errro in cacahing", err)
	}
	// range through kine of *-page.html
	for _, page := range pages {
		// page
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return Cacahe, err
		}
		//base page
		matche, err := filepath.Glob("/home/ihtgoot/web_dev/Section_3/templates/*-layout.html")
		if err != nil {
			return Cacahe, err
		}
		// use base
		if len(matche) > 0 {
			templateSet, err = templateSet.ParseGlob("/home/ihtgoot/web_dev/Section_3/templates/*layout.html")
			if err != nil {
				return Cacahe, err
			}
		}
		// adding the page to cacahe
		Cacahe[name] = templateSet
	}
	return Cacahe, nil
}
