package libs

import (
	"html/template"
	"net/http"
)

func NewTemplateCache() *TemplateCache {
	return &TemplateCache{make(map[string]*template.Template)}
}

type TemplateCache struct {
	cache map[string]*template.Template
}

func (t *TemplateCache) HasCache(templateFile string) bool {
	_, ok := t.cache[templateFile]
	return ok
}

func (t *TemplateCache) ParseFiles(templateFile string) (*template.Template, error) {
	tc, err := template.ParseFiles(templateFile)
	if err != nil {
		t.cache[templateFile] = tc
	}
	return tc, err
}

func (t *TemplateCache) Get(templateFile string) (*template.Template, error) {
	tc, ok := t.cache[templateFile]
	if ok {
		return tc, nil
	} else {
		tc, err := t.ParseFiles(templateFile)
		return tc, err
	}
}

var templateCache *TemplateCache = nil

func Viewer(templateFile string) func(http.ResponseWriter, interface{}) {
	if templateCache == nil {
		templateCache = NewTemplateCache()
	}
	tc, err := templateCache.Get(templateFile)
	if err != nil {
		panic(err)
	}
	view := func(w http.ResponseWriter, data interface{}) {
		tc.Execute(w, data)
	}
	return view
}
