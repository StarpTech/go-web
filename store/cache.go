package store

type Cache interface {
	Ping() error
}
