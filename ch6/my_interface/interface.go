package myinterface

type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) error
}

type Closer interface {
	Close() error
}
