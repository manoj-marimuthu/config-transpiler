package parser

import (
	"config-transpiler/config"
	"gopkg.in/yaml.v3"
	"os"
)

func Parse(filename string) config.CloudInit{
	data,err := os.ReadFile(filename)
	if err != nil{
		panic(err)
	}
	var cfg config.CloudInit
	yaml.Unmarshal(data,&cfg)
	return cfg
}


