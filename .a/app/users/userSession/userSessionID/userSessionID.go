package userSessionID

import "imageresizerservice/library/id"

type UserSessionID string

func Gen() UserSessionID {
	return UserSessionID(id.Gen())
}
