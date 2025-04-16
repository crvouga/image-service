package auth

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"net/http"
)

// IsLoggedIn checks if the user is logged in by checking if they have a valid user session
func IsLoggedIn(appCtx *appCtx.AppCtx, r *http.Request) bool {
	reqCtxInst := reqCtx.FromHttpRequest(appCtx, r)
	return reqCtxInst.UserSession != nil
}
