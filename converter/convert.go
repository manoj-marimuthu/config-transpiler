package converter

import (
	"config-transpiler/config"
)

type converter func(in config.CloudConfig,out *config.Butane)

var converters []converter

func register(f converter){
	converters = append(converters,f)
}

func Convert(in config.CloudConfig) config.Butane{
	out := config.Butane{}
	for _,converter := range converters{
		converter(in,&out)
	}
	return out
}
