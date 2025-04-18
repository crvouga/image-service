package page

import (
	"html/template"
	"net/http"
)

func Respond(pageData any, templatePaths ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Always include base templates
		allTemplatePaths := []string{
			"./app/ui/page/page.html",
			"./app/ui/breadcrumbs/breadcrumbs.html",
			"./app/ui/icons.html",
			"./app/ui/header.html",
			"./app/ui/pageHeader/pageHeader.html",
			"./app/ui/mainMenu/mainMenu.html",
		}

		// Add any additional template paths
		allTemplatePaths = append(allTemplatePaths, templatePaths...)

		// Define function map for templates
		funcMap := template.FuncMap{}

		// Create template with function map
		tmpl, err := template.New("page.html").Funcs(funcMap).ParseFiles(allTemplatePaths...)
		if err != nil {
			errStr := err.Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		// Set content type header
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Execute the root template which will include all others
		if err := tmpl.ExecuteTemplate(w, "page.html", pageData); err != nil {
			errStr := err.Error()
			http.Error(w, errStr, http.StatusInternalServerError)
		}
	}
}
