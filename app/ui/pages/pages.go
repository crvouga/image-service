package pages

import (
	"imageService/app/ui/confirmationPage"
	"imageService/app/ui/errorPage"
	"imageService/app/ui/notFoundPage"
	"imageService/app/ui/successPage"
	"net/http"
)

func Router(mux *http.ServeMux) {
	confirmationPage.Router(mux)
	errorPage.Router(mux)
	successPage.Router(mux)
	notFoundPage.Router(mux)
}
