package config

type User struct{
	Name string `yaml:"name"`
	Gecos string `yaml:"gecos"`
}

type CloudConfig struct{
	Users []User `yaml:"users"`
}

