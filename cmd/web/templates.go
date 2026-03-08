package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/sumanththota/snippetbox/pkg/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	//Initlaize a new map to act as the cache

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		//opens the base page and parses it into a template object of type *template.Template
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}
		// 		ts.templates = {
		//    "home.page.html": TemplateObject
		//    "base.layout.html": TemplateObject
		//    "admin.layout.html": TemplateObject
		// } this is how the ts looks inside
		cache[name] = ts
	}
	return cache, nil
}
