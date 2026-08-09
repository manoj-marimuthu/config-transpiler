package converter

import (
	"config-transpiler/config"
)


func groupsConvert(in config.CloudConfig, out *config.Butane){
	for _, group := range in.Groups{
		out.Passwd.Groups = append(out.Passwd.Groups,group)
	}
}

func init(){
	register(groupsConvert)
}
