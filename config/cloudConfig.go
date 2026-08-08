package config

type CloudConfigUser struct{
	Name string `yaml:"name"`
	Gecos string `yaml:"gecos"`
	Groups string `yaml:"groups"`
	Shell string `yaml: "shell"`
}

type CloudConfig struct{
	Users []CloudConfigUser `yaml:"users"`
}

