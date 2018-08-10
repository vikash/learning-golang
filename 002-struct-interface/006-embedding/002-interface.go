package main

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWriter interface {
	Reader
	Writer
}
