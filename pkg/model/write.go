package model

import (
	"io"
	"sync"
)

// Writer is the interface that uses for writing data to output
type Writer interface {
	io.Writer
	io.Closer
}
type HashWriter interface {
	WriteToHAsh() (*sync.Map, error)
}
