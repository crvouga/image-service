package auth

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/ctx/reqCtx"
	"net/http"
)

// IsLoggedIn checks if the user is logged in by checking if they have a valid user session
func IsLoggedIn(appCtx *appContext.AppCtx, r *http.Request) bool {
	reqCtxInst := reqCtx.FromHttpRequest(appCtx, r)
	return reqCtxInst.UserSession != nil
}
