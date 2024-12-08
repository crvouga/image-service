package keyValueDb

type KeyValueDb interface {
	Get(key string) (string, error)
	Put(value string) error
	Zap(key string) error
}
