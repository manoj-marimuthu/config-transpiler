package converter

import (
	"config-transpiler/config"
)

func  userConvert(in config.CloudInit,out *config.Butane){
	for _,user := range in.Users{
		butaneUser := config.User{
			Name : user.Name,
			Gecos : user.Gecos,
		}
		out.Passwd.Users = append(out.Passwd.Users, butaneUser)
	}
}

func init(){
	register(userConvert)
}
