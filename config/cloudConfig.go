package config

type CloudConfigUser struct {
	Name   string `yaml:"name"`
	Gecos  string `yaml:"gecos"`
	Groups string `yaml:"groups"`
	Shell  string `yaml: "shell"`
}

type CloudConfig struct {
	Groups []string `yaml:"groups"`
	Users []CloudConfigUser `yaml:"users"`
}
