package view

import (
	"crawder/single/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	Template *template.Template
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.Template.Execute(w, data)
}

func CreateView(filename string) SearchResultView {
	return SearchResultView{
		Template: template.Must(template.ParseFiles(filename)),
	}
}
