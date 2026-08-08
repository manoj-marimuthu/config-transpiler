package config

type ButaneUser struct{
	Name string `yaml:"name"`
	Gecos string `yaml:"gecos,omitempty"`
	Groups []string `yaml:"groups,omitempty"`
	Shell string `yaml:"shell,omitempty"`
}

type Passwd struct{
	Users []ButaneUser `yaml:"users"`	
}

type Butane struct{
	Passwd Passwd `yaml:"passwd"`
}
