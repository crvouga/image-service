package pages

import (
	"imageresizerservice/app/ui/confirmationPage"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/ui/notFoundPage"
	"imageresizerservice/app/ui/successPage"
	"net/http"
)

func Router(mux *http.ServeMux) {
	confirmationPage.Router(mux)
	errorPage.Router(mux)
	successPage.Router(mux)
	notFoundPage.Router(mux)
}
