package config


type Passwd struct{
	Users []User `yaml:"users"`	
}

type Butane struct{
	Passwd Passwd `yaml:"passwd"`
}
