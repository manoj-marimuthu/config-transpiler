package writer

import (
	"config-transpiler/config"
	"gopkg.in/yaml.v3"
	"os"
)

func Write(bcfg config.Butane, outfilename string) {
	data, err := yaml.Marshal(bcfg)
	if err != nil {
		panic(err)
	}
	errorFromFileWrite := os.WriteFile(outfilename, data, 0644)
	if errorFromFileWrite != nil {
		panic(errorFromFileWrite)
	}
}
