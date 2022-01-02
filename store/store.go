package store

type Store interface {
	Get(key string) (string, error)
	Set(key, value string) error
	SetBytes(key, value []byte) error
	Close() error
}
