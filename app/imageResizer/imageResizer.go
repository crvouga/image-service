package imageResizer

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/imageResizer/imageResizerPage"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	imageResizerPage.Router(mux, appCtx)
}
