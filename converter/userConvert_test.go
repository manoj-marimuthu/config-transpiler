package converter

import (
	"config-transpiler/config"
	"reflect"
	"testing"
)

func TestUserConverter(t *testing.T) {
	tests := []struct {
		in  config.CloudConfig
		out config.Butane
	}{
		// test case 1
		{
			config.CloudConfig{
				Users: []config.CloudConfigUser{
					{Name: "example", Gecos: "example-Gecos", Shell: "/bin/bash", Uid: 1004, HomeDir: "/home/exampleHome"},
					{Name: "example2", Gecos: "example-2-Gecos", Shell: "/bin/bash", Uid: 1023, HomeDir: "/home/example2Home"},
				},
			},
			config.Butane{
				Passwd: config.Passwd{
					Users: []config.ButaneUser{
						{Name: "example", Gecos: "example-Gecos", Shell: "/bin/bash", Uid: 1004, Home_Dir: "/home/exampleHome"},
						{Name: "example2", Gecos: "example-2-Gecos", Shell: "/bin/bash", Uid: 1023, Home_Dir: "/home/example2Home"},
					},
				},
			},
		},
		// test case 2
		{
			config.CloudConfig{
				Users: []config.CloudConfigUser{
					{Name: "example3", Gecos: "", Shell: "", Uid: 2300, HomeDir: "/home/example3Home"},
				},
			},
			config.Butane{
				Passwd: config.Passwd{
					Users: []config.ButaneUser{
						{Name: "example3", Gecos: "", Shell: "", Uid: 2300, Home_Dir: "/home/example3Home"},
					},
				},
			},
		},
		// test case 3
		{
			config.CloudConfig{},
			config.Butane{},
		},
	}
	for i, test := range tests {
		var out config.Butane
		userConvert(test.in, &out)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test %d, Expected: %+v Received: %+v \n \n", i, test.out, out)
		}
	}
}
