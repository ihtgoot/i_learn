// save our cache in golabally configurable file
// importred everywhere in the web app
package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig is a struct holding die application's configuration
type AppConfig struct {
	TemplateCacahe map[string]*template.Template
	UseCache       bool
	InfoLog        *log.Logger // handel errro centrally , write to a log file the errors
	InProduction   bool
	Session        *scs.SessionManager
}
