package render

import (
	"bytes"
	"github.com/davidxchen/bookings/pkg/config"
	"github.com/davidxchen/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// Template rendering html templates
func Template(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache form the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache creating template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// get all files named *.page.gohtml from ./templates
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return templateCache, err
	}

	// range through all files ending with *.page.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
		}

		templateCache[name] = ts
	}

	return templateCache, nil
}
