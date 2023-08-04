package pkg

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the config
type Config struct {
	GetVgangDataApiUrl     string
	OutputFilePath         string
	DataGetterWorkerSize   int
	DataRendererWorkerSize int
	OutputFormat           string
}

// NewConfig creates a new config
func NewConfig() *Config {
	v := viper.New()
	//config file name
	v.SetConfigFile("./config.yml")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	// v.AddConfigPath("../config")

	// v.SetConfigName("config")
	//config file path
	// setDefaultValues(v)

	//for use env variables
	v.AutomaticEnv()

	//type of config file
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// fmt.Println(v.GetInt("data_getter_worker_size"))
	return &Config{
		GetVgangDataApiUrl:     v.GetString("get_vgang_data_api_url"),
		OutputFormat:           v.GetString("output_format"),
		OutputFilePath:         v.GetString("output_file_path"),
		DataGetterWorkerSize:   v.GetInt("data_getter_worker_size"),
		DataRendererWorkerSize: v.GetInt("data_renderer_worker_size"),
	}
}
