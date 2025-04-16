package useLinkAction

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/login/link"
)

type Fixture struct {
	AppCtx       appCtx.AppCtx
	ExistingLink link.Link
}

func NewFixture() Fixture {

	f := Fixture{AppCtx: appCtx.NewTest()}

	return f
}
