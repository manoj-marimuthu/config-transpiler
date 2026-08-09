package parser

import (
	"config-transpiler/config"
	"gopkg.in/yaml.v3"
	"os"
)

func Parse(filename string) config.CloudConfig {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var cfg config.CloudConfig
	yaml.Unmarshal(data, &cfg)
	return cfg
}
