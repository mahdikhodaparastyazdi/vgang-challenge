package writer

import (
	"log"
	"os"

	"github.com/mahdikhodaparast/vgang-challenge/pkg"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
)

// fileResponseWriter writes the response to a file
type fileResponseWriter struct {
	resultFile *os.File
	config     *pkg.Config
}

// NewFileResponseWriter creates a new fileResponseWriter
func NewFileResponseWriter(config *pkg.Config) model.Writer {
	resultFile, err := os.Create(config.OutputFilePath)
	if err != nil {
		log.Fatal("can not create result file", err)
	}
	return &fileResponseWriter{
		config:     config,
		resultFile: resultFile,
	}
}

// Close closes the file
func (f fileResponseWriter) Close() error {
	err := f.resultFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func (f fileResponseWriter) Write(p []byte) (n int, err error) {
	_, err = f.resultFile.WriteString(string(p))
	if err != nil {
		return 0, err
	}
	return 1, err
}
