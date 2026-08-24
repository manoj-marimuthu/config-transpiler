package config

type CloudConfigUser struct {
	Name   string `yaml:"name"`
	Gecos  string `yaml:"gecos"`
	Groups string `yaml:"groups"`
	Shell  string `yaml: "shell"`
	Uid int `yaml: "uid"`
}

type CloudConfig struct {
	Groups []string `yaml:"groups"`
	Users []CloudConfigUser `yaml:"users"`
	RunCmd []string `yaml:"runcmd"`
}
