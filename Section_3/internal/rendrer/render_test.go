package rendrer

import (
	"log"
	"net/http"
	"testing"

	"github.com/ihtgoot/i_learn/Section_3/internal/models"
)

func TestAddDefaultData(t *testing.T) {

	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Fatal(err)
	}

	session.Put(r.Context(), "flash", "a flash message")
	session.Put(r.Context(), "error", "a error message")
	session.Put(r.Context(), "warning", "a warning message")

	result := AddDefaultData(&td, r)
	if result.Flash != "a flash message" {
		t.Error("flash message not found ::: code fat gaya")
	} else {
		t.Log("OK Flash")
	}
	if result.Error != "a error message" {
		t.Error("error message not found ::: code fat gaya")
	} else {
		t.Log("OK Error")
	}
	if result.Warning != "a warning message" {
		t.Error("warning message not found ::: code fat gaya")
	} else {
		t.Log("OK Warning")
	}

}

func TestRenderTemplate(t *testing.T) {

	// var pathToTemplate = "../../templates/"

	tc, err := CreateTempleteCache()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("template cache keys:")
	for k := range tc {
		t.Log("-", k)
	}

	app.TemplateCacahe = tc
	app.UseCache = true

	r, err := getSession()
	if err != nil {
		log.Fatal(err)
	}

	var ww myWriter

	key := "about-page.tpml"
	if _, ok := tc[key]; !ok {
		t.Fatalf("template %s not found in cache; check printed cache keys", key)
	}

	err = RenderTemplate(&ww, r, "home-page.tpml", &models.TemplateData{})
	if err != nil {
		t.Error("RenderTemplate did not work , writing template to browser failed : ", err)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/dummy-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestNewTemplate(t *testing.T) {
	// test new templated can be called or not
	NewTemplate((app))
}

func TestCreateTemplateCache(t *testing.T) {
	//var path/ToTemplate = "../../templates/"

	_, err := CreateTempleteCache()
	if err != nil {
		t.Error(err)
	}
}
