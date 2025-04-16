package useLinkAction

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/users/login/link"
)

type Fixture struct {
	AppCtx       appContext.AppCtx
	ExistingLink link.Link
}

func NewFixture() Fixture {

	f := Fixture{AppCtx: appContext.NewTest()}

	return f
}
