package models

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
}
