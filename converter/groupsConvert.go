package converter

import (
	"config-transpiler/config"
)

func groupsConvert(in config.CloudConfig, out *config.Butane) {
	for _, group := range in.Groups {
		butanePasswdGroups := config.ButanePasswdGroups{
			Name: group,
		}
		out.Passwd.Groups = append(out.Passwd.Groups, butanePasswdGroups)
	}
}

func init() {
	register(groupsConvert)
}
