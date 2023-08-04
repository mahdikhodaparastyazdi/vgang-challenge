package cmd

import (
	"os"

	"github.com/mahdikhodaparast/vgang-challenge/pkg"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/model"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/repository"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/service"
	"github.com/mahdikhodaparast/vgang-challenge/pkg/service/writer"
)

// Factory creates the service  dependencies
type Factory struct {
	Config         *pkg.Config
	PipelineStages []model.Runner
	DoneChannel    chan bool
	Pipeline       model.HashWriter
}

// NewFactory creates a new Factory
func NewFactory() (*Factory, error) {
	config := pkg.NewConfig()
	done := make(chan bool)
	productURlChan := make(chan [2]int, 1000)
	productIDChan := make(chan int, 1000)

	//Initialize the repositories
	vgangRepository, err := repository.NewVgangRepository(config)
	if err != nil {
		return nil, err
	}
	//Get Token for apis
	getToken, err := vgangRepository.GetToken()
	if err != nil {
		return nil, err
	}
	//Initialize the services
	writerService := getWriterService(config)

	// Initialize the pipeline stages
	generatorService, err := service.NewGeneratorService(config, vgangRepository, productURlChan, getToken)
	if err != nil {
		return nil, err
	}
	dataGetterService, err := service.NewDataGetterService(config, vgangRepository, productURlChan, productIDChan, getToken)
	if err != nil {
		return nil, err
	}
	rendererService, err := service.NewDataRendererService(config, productIDChan, writerService, done)
	if err != nil {
		return nil, err
	}

	//Register the pipeline stages
	pipelineStages := registerPipelineStages(generatorService, dataGetterService, rendererService)

	hashWriter := writer.NewMapResponseWriter(config)

	return &Factory{
		Config:         config,
		PipelineStages: pipelineStages,
		DoneChannel:    done,
		Pipeline:       hashWriter,
	}, nil
}

func getWriterService(config *pkg.Config) model.Writer {
	switch config.OutputFormat {
	case "file":
		return writer.NewFileResponseWriter(config)
	case "stdout":
		return writer.NewStdResponseWriter(config, os.Stdout)
	default:
		return writer.NewStdResponseWriter(config, os.Stdout)
	}
}

func registerPipelineStages(stages ...model.Runner) []model.Runner {
	pipeLineStages := append([]model.Runner{}, stages...)
	return pipeLineStages
}
