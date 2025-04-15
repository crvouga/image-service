package useLinkAction

import (
	"imageresizerservice/app/ctx/appCtx"
)

type Fixture struct {
	AppCtx appCtx.AppCtx
}

func NewFixture() Fixture {
	f := Fixture{AppCtx: appCtx.NewTest()}
	return f
}
