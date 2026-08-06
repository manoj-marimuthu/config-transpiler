package config

type User struct{
	Name string `yaml:"name"`
	Gecos string `yaml:"gecos"`
}

type CloudInit struct{
	Users []User `yaml:"users"`
}

