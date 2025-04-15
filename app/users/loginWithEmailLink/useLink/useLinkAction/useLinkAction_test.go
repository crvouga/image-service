package useLinkAction

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/loginWithEmailLink/link"
)

type Fixture struct {
	AppCtx       appCtx.AppCtx
	ExistingLink link.Link
}

func NewFixture() Fixture {

	f := Fixture{AppCtx: appCtx.NewTest()}

	return f
}
