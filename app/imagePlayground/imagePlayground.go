package imagePlayground

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/imagePlayground/imagePlaygroundPage"

	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	imagePlaygroundPage.Router(mux, ac)
}
