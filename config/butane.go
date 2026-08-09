package config

type ButaneUser struct {
	Name   string   `yaml:"name"`
	Gecos  string   `yaml:"gecos,omitempty"`
	Groups []string `yaml:"groups,omitempty"`
	Shell  string   `yaml:"shell,omitempty"`
}

type Passwd struct {
	Users []ButaneUser `yaml:"users"`
	Groups []string `yaml:"groups"` 
}

type Butane struct {
	Variant string `yaml:"variant"`
	Version string `yaml:"version"`
	Passwd  Passwd `yaml:"passwd"`
}
