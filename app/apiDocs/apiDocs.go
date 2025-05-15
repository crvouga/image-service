package apiDocs

import (
	"imageService/app/apiDocs/apiDocsPage"
	"imageService/app/ctx/appCtx"

	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	apiDocsPage.Router(mux, ac)
}
