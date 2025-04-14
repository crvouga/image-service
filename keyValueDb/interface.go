package keyValueDb

import "imageresizerservice/uow"

type KeyValueDb interface {
	Get(key string) (string, error)
	Put(uow *uow.Uow, key string, value string) error
	Zap(uow *uow.Uow, key string) error
}
