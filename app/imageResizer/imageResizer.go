package imageResizer

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/imageResizer/imageResizerPage"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appContext.AppCtx) {
	imageResizerPage.Router(mux, appCtx)
}
