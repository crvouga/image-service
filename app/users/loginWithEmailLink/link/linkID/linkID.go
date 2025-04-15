package linkID

import "imageresizerservice/library/id"

type LinkID string

func Gen() LinkID {
	return LinkID(id.Gen())
}

func New(id string) LinkID {
	return LinkID(id)
}
