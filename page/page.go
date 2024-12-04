package page

import (
	"html/template"
	"net/http"
)

func Handler(pageTemplatePath string, pageData any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./page/page.html", pageTemplatePath)

		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, pageData); err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}
