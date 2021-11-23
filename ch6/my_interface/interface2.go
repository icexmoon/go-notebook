package myinterface

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	ReadWriter
	Closer
}

type ReadWriteClosePrinter interface {
	ReadWriteCloser
	Print()
}
