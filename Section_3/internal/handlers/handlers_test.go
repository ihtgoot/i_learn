package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

type handlerTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}

var allTheHandlerTests = handlerTests{
	{"not-existing-route", "/not-existing-dummy-route", "GET", []postData{}, http.StatusNotFound},
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"eremite", "/eremite", "GET", []postData{}, http.StatusOK},
	{"couple", "/couple", "GET", []postData{}, http.StatusOK},
	{"family", "/family", "GET", []postData{}, http.StatusOK},
	{"reservation", "/reservation", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-reservation", "/reservation", "POST", []postData{
		{key: "startingEND", value: "23-03-2025"},
		{key: "endingEND", value: "23-03-2026"},
	}, http.StatusOK},
	{"post-reservation-json", "/reservation", "POST", []postData{
		{key: "startingEND", value: "23-03-2025"},
		{key: "endingEND", value: "23-03-2026"},
	}, http.StatusOK},
	{"post-make-reservation", "/make-reservation", "POST", []postData{
		{key: "full_name", value: "john green"},
		{key: "phone", value: "484945616416"},
		{key: "email", value: "efubghh@idfjbg.ijefbg"},
	}, http.StatusOK},
}

func TestAllTheHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()
	for _, test := range allTheHandlerTests {
		if test.method == "GET" {
			response, err := testServer.Client().Get(testServer.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			} else {
				t.Logf("%s 		OK", test.name)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected status code %d but got %d NOT OK", test.name, test.expectedStatusCode, response.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, param := range test.params {
				t.Logf("%s : %s", param.key, param.value)
				values.Add(param.key, param.value)
			}
			response, err := testServer.Client().PostForm(testServer.URL+test.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			} else {
				t.Logf("%s 		OK", test.name)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected status code %d but got %d NOT OK", test.name, test.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
