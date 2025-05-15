package userSessionID

import "imageService/library/id"

type UserSessionID string

func Gen() UserSessionID {
	return UserSessionID(id.Gen())
}
