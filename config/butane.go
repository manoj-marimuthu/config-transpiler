package config

type ButaneUser struct {
	Name   string   `yaml:"name"`
	Gecos  string   `yaml:"gecos,omitempty"`
	Groups []string `yaml:"groups,omitempty"`
	Shell  string   `yaml:"shell,omitempty"`
	Uid    int      `yaml:"uid,omitempty"`
}

type ButanePasswdGroups struct {
	Name string `yaml:"name"`
}

type Passwd struct {
	Users  []ButaneUser         `yaml:"users"`
	Groups []ButanePasswdGroups `yaml:"groups"`
}

type ButaneSystemdUnit struct {
	Name     string `yaml:"name"`
	Enabled  bool   `yaml:"enabled"`
	Contents string `yaml:"contents"`
}

type Systemd struct {
	Units []ButaneSystemdUnit `yaml:"units"`
}

type FileContents struct {
	Inline string `yaml:"inline"`
}

type FileUser struct {
	Name string `yaml:"name"`
}

type FileGroup struct {
	Name string `yaml:"name"`
}

type File struct {
	Path     string       `yaml:"path"`
	Mode     int          `yaml:"mode"`
	Contents FileContents `yaml:"contents"`
	User     FileUser     `yaml:"user"`
	Group    FileGroup    `yaml:"group"`
}

type Storage struct {
	Files []File `yaml:"files"`
}

type Butane struct {
	Variant string  `yaml:"variant"`
	Version string  `yaml:"version"`
	Passwd  Passwd  `yaml:"passwd"`
	Systemd Systemd `yaml:"systemd"`
	Storage Storage `yaml:"storage"`
}
