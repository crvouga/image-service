package apiDocs

import (
	"imageresizerservice/app/apiDocs/apiDocsPage"
	"imageresizerservice/app/ctx/appCtx"

	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	apiDocsPage.Router(mux, ac)
}
