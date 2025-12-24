package form

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Errors map[string][]string

// form is a type type holding a general form structure includeing an url.Value object
type Forms struct {
	url.Values
	Errors Errors
}

// new is a function to initialize form structure
func New(data url.Values) *Forms {
	return &Forms{
		Values: data,
		Errors: make(Errors),
	}
}

// has check for the existance of a form field in the post and esusre it is not empty
func (f *Forms) Has(field string) bool {
	formfield := f.Get(field)
	if formfield == "" {
		//f.Errors[field] = append(f.Errors[field], "this cannot be empty")
		return false
	}
	return true
}

// valid false in case of errors otherwise true
func (f *Forms) Valid() bool {
	return len(f.Errors) == 0
}

// required checks for the existance of form in form field in the post and ensure they are not empty
func (f *Forms) Required(fields ...string) bool {
	var valid bool
	valid = true
	for _, field := range fields {
		formfield := f.Get(field)
		if formfield == "" {
			f.Errors[field] = append(f.Errors[field], "this cannot be empty")
			return false
		}
	}
	return valid
}

// Get returns the first error message for the specified field from the Errors map.
// If there are no errors for the field, it returns an empty string.
// Useful for rendering form validation errors in templates (shows the first error for a field).
func (e Errors) Get(field string) string {
	if len(e[field]) > 0 {
		return e[field][0]
	}
	return ""
}

// minlength returns false if the field value is shorter than a given length , other wise true
func (f *Forms) MinLength(field string, length int) bool {
	actualLength := f.Get(field)
	if len(strings.TrimSpace(actualLength)) < length {
		f.Errors[field] = append(f.Errors[field], fmt.Sprint("this cannot be samller than "+strconv.Itoa(length)+" characters"))
		return false
	}
	return true
}

// isEmail checks if value of field is a valid emil address
func (f *Forms) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors[field] = append(f.Errors[field], fmt.Sprint("this is not a valid email"))
		return false
	}
	return true
}
