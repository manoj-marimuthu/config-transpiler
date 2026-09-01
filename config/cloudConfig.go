package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type CloudConfigUser struct {
	Name    string `yaml:"name"`
	Gecos   string `yaml:"gecos"`
	Groups  string `yaml:"groups"`
	Shell   string `yaml: "shell"`
	Uid     int    `yaml: "uid"`
	HomeDir string `yaml: "homedir"`
	System  bool   `yaml: "system"`
}

type CloudConfigWriteFiles struct {
	Path        string `yaml:"path"`
	Content     string `yaml:"content"`
	Permissions string `yaml:"permissions"`
	Owner       string `yaml:"owner"`
}

type CloudConfigRunCmd struct {
	Command string
	Args []string
}

func (r* CloudConfigRunCmd) UnmarshalYAML(value* yaml.Node) error{
	switch value.Kind{
		case yaml.ScalarNode:
			r.Command = value.Value
			break
		case yaml.SequenceNode:
			for _, arg := range value.Content{
				r.Args = append(r.Args, arg.Value)
			}
			break
		default:
			return fmt.Errorf("runcmd requires a string or a list")
	}
	return nil
}
type CloudConfig struct {
	Groups     []string                `yaml:"groups"`
	Users      []CloudConfigUser       `yaml:"users"`
	RunCmd     []CloudConfigRunCmd     `yaml:"runcmd"`
	WriteFiles []CloudConfigWriteFiles `yaml:"write_files"`
}
