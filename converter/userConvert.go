package converter

import (
	"config-transpiler/config"
	"strings"
)

func userConvert(in config.CloudConfig, out *config.Butane) {
	for _, user := range in.Users {
		var groups []string
		if user.Groups != "" {
			groups = strings.Split(user.Groups, ",")
			for i, group := range groups {
				groups[i] = strings.TrimSpace(group)
			}
		}
		butaneUser := config.ButaneUser{
			Name:   user.Name,
			Gecos:  user.Gecos,
			Groups: groups,
			Shell:  strings.TrimSpace(user.Shell),
			Uid:    user.Uid,
		}
		out.Passwd.Users = append(out.Passwd.Users, butaneUser)
	}
}

func init() {
	register(userConvert)
}
