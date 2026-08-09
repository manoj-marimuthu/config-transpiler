package main

import (
	"config-transpiler/config"
	"config-transpiler/converter"
	"config-transpiler/parser"
	"config-transpiler/writer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input and output yaml file names")
		return
	}
	var infilename string = os.Args[1]
	var outfilename string = os.Args[2]
	var cfg config.CloudConfig = parser.Parse(infilename)
	bcfg := converter.Convert(cfg)
	writer.Write(bcfg, outfilename)
}
