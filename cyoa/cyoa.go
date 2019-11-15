package cyoa

import (
	"html/template"
	"log"
	"net/http"
)

// Story is a map to map the string values ttheiro  respective data
type Story map[string]book

//Rendertemplate is a renderfunction to be exported so that the template gets exported.
func Rendertemplate(w http.ResponseWriter, templatefile string, b book) {
	t, err := template.ParseFiles(templatefile)
	if err != nil {
		log.Fatal("error encountered while parsing the template:", err)
	}
	t.Execute(w, b)
}

type book struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"paragraphs"`
	Options   []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
