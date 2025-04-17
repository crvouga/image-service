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

		tmpl, err := template.ParseFiles(allTemplatePaths...)
		if err != nil {
			errStr := err.Error()
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, pageData); err != nil {
			errStr := err.Error()
			http.Error(w, errStr, http.StatusInternalServerError)
		}
	}
}
