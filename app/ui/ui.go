package ui

import (
	"imageresizerservice/app/ui/confirmationPage"
	"imageresizerservice/app/ui/resultPage"
	"net/http"
)

func Router(mux *http.ServeMux) {
	confirmationPage.Router(mux)
	resultPage.Router(mux)
}
