package auth

import (
	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"net/http"
)

// IsLoggedIn checks if the user is logged in by checking if they have a valid user session
func IsLoggedIn(ac *appCtx.AppCtx, r *http.Request) bool {
	rc := reqCtx.FromHttpRequest(ac, r)
	return rc.UserSession != nil
}
