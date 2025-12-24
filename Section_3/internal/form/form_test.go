package form

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/an-url", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	// test form is valid or not
	// valid false in case of errors otherwise true
	if !isValid {
		t.Error("expected valid got invalid")
	}
}

// test if required checks for the existance of form in form field in the post and ensure they are not empty
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/an-url", nil)
	form := New(r.PostForm)

	form.Required("some_thing", "some_other_thing", "an_other_thing")
	if form.Valid() {
		t.Error("expected invalid got valid")
	}

	postData := url.Values{}
	postData.Add("a", "a value")
	postData.Add("b", "b value")
	postData.Add("c", "c value")
	r = httptest.NewRequest("POST", "/an-url", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("expected valid got invalid")
	}
}

// has check for the existance of a form field in the post and esusre it is not empty
func TestForm_Has(t *testing.T) {

	postedData := url.Values{}
	form := New(postedData)
	has := form.Has("whatever")
	if has {
		t.Error("reported existence but deosnot exist")
	}

	postedData = url.Values{}
	postedData.Add("Whatever", "whatsoever")
	form = New(postedData)
	has = form.Has("Whatever")
	if !has {
		t.Error("reported not existence but deos exist")
	}

}

func TestForm_MinLength(t *testing.T) {
	length := 3
	postedData := url.Values{}
	var testv string
	form := New(postedData)

	form.MinLength("testv", length)
	if form.Valid() {
		t.Error("forms shows min length for non existinig data")
	}
	isError := form.Errors.Get("testv")
	if isError == "" {
		t.Error("form repost na erro but doesnot exist")
	}

	testiv := "hi"
	postedData.Add("testiv", testiv)
	form.MinLength("testiv", length)
	if form.Valid() {
		t.Error("forms shows min length met but is shorter")
	}
	isError = form.Errors.Get("testiv")
	if isError == "" {
		t.Error("form repost no error but error exist")
	}

	testv = "hi_there"
	postedData.Add("testv", testv)
	form.MinLength("testv", length)
	if form.Valid() {
		t.Error("forms shows min length met but is is invalid")
	}
	isError = form.Errors.Get("testv")
	if isError == "" {
		t.Error("form repost no error but error exist")
	}

}

func TestForm_IsEmail(t *testing.T) {
	emailv := "hdvgf@gmail.com"
	emailiv := "dofjgbndfjgn"

	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("returned valid in no existance field")
	}

	postedData = url.Values{}
	postedData.Add("v", emailv)
	form = New(postedData)
	form.IsEmail("v")
	if !form.Valid() {
		t.Error("returned invalid in valid email")
	}

	postedData = url.Values{}
	postedData.Add("iv", emailiv)
	form = New(postedData)
	form.IsEmail("iv")
	if form.Valid() {
		t.Error("returned valid in invalid email")
	}
}
