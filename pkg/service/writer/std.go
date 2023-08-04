package writer

import (
	"fmt"
	"io"

	"github.com/mahdikhodaparast/vgang-challenge/pkg"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
)

// stdResponseWriter is a writer that writes to stdout
type stdResponseWriter struct {
	config *pkg.Config
	writer io.Writer
}

// Write on stdout
func (s *stdResponseWriter) Write(p []byte) (n int, err error) {
	//TODO implement me
	_, err = fmt.Fprintf(s.writer, string(p))
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// NewStdResponseWriter creates a new stdResponseWriter
func NewStdResponseWriter(config *pkg.Config, writer io.Writer) model.Writer {
	return &stdResponseWriter{
		config: config,
		writer: writer,
	}
}

// Close writer interface
func (s *stdResponseWriter) Close() error {
	return nil
}
