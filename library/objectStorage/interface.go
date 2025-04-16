package objectStorage

type ObjectStorage interface {
	Get(path string) ([]byte, error)
	Put(path string, data []byte) error
	Exists(path string) (bool, error)
	Zap(path string) error
}
