package log

type Writable interface {
	Write(log Log) error
	Close() error
}

type Readable interface {
	Read(start, end int64) ([]string, error)
	Close() error
}
