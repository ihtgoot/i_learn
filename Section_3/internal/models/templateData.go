package models

import "github.com/ihtgoot/i_learn/Section_3/internal/form"

// templateData holds any kind of data sent from handler to template
type TemplateData struct {
	StringMap map[string]string
	IntString map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *form.Forms
}
