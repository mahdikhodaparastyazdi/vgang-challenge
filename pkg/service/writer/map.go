package writer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/mahdikhodaparast/vgang-challenge/pkg"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
)

type mapResponseWriter struct {
	config *pkg.Config
}

func NewMapResponseWriter(config *pkg.Config) model.HashWriter {

	return &mapResponseWriter{
		config: config,
	}
}

func (*mapResponseWriter) WriteToHAsh() (*sync.Map, error) {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatal("can not create result file", err)
	}

	// Construct the path to result.txt
	resultFilePath := filepath.Join(projectRoot, "result.txt")

	// Read the contents of result.txt
	data, err := ioutil.ReadFile(resultFilePath)
	if err != nil {
		log.Fatal("can not create result file", err)
	}
	fmt.Print(data)
	return nil, nil
}
